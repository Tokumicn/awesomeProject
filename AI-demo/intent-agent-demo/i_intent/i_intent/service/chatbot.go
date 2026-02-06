package service

import (
	"context"
	"fmt"
	"i_intent/data/model"
	"i_intent/llm"
	"i_intent/llm/prompts"
	"i_intent/utils"
	"log"
	"strings"
	"sync"

	"github.com/gogf/gf/v2/frame/g"
)

type ChatbotService struct {
	sceneTemplates      map[string]models.SceneConfig
	currentPurpose      string
	lastRecognizedScene string
	processors          map[string]SceneProcessor
	sceneSlots          map[string][]models.Slot
	chatHistory         []models.ChatMessage
	isSlotFilling       bool
	mu                  sync.RWMutex
}

var _chatbotService *ChatbotService

func NewChatbotService(sceneTemplates map[string]models.SceneConfig) *ChatbotService {
	if _chatbotService == nil {
		_chatbotService = &ChatbotService{
			sceneTemplates: sceneTemplates,
			processors:     make(map[string]SceneProcessor),
			sceneSlots:     make(map[string][]models.Slot),
			chatHistory:    []models.ChatMessage{},
			isSlotFilling:  false,
		}
	}
	return _chatbotService
}

func (cs *ChatbotService) ProcessMultiQuestion(ctx context.Context, userInput string) string {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	cs.chatHistory = append(cs.chatHistory, models.ChatMessage{
		Role:    "user",
		Content: userInput,
	})

	// 首次场景 =============================START
	if cs.currentPurpose == "" {
		// 无意图场景  开始初次识别
		cs.recognizeIntent(ctx, userInput)
	}

	if cs.currentPurpose == "" {
		// 未能获取到意图场景 构建无意图回复
		response := cs.generateNoSceneResponse(ctx, userInput)
		cs.chatHistory = append(cs.chatHistory, models.ChatMessage{
			Role:    "assistant",
			Content: response,
		})
		return response
	}
	// 首次场景 ============================= END

	// 多轮 =============================START
	// 以下是识别到意图 开始词槽填充
	cs.isSlotFilling = true
	g.Log().Printf(ctx, "进入词槽填充 当前场景[current_purpose: %s]", cs.currentPurpose)

	// 检查是否转换了意图场景
	if cs.detectSceneSwitch(ctx, userInput) {
		// 转换了意图场景则清理之前的意图
		cs.clearCurrentScene(ctx)
		// 多轮对话  上轮的意图已经变更 因此需要重新识别意图
		cs.recognizeIntent(ctx, userInput)

		if cs.currentPurpose != "" {
			processor, err := cs.getProcessorForScene(ctx, cs.currentPurpose)
			if err != nil {
				return "抱歉，无法找到场景处理器。"
			}

			response, err := processor.Process(ctx, userInput, cs.chatHistory)
			if err != nil {
				response = fmt.Sprintf("处理出错: %v", err)
			}

			if !strings.HasPrefix(response, "请问") && !strings.HasPrefix(response, "抱歉，无法找到场景") {
				cs.clearCurrentScene(ctx)
			}

			cs.chatHistory = append(cs.chatHistory, models.ChatMessage{
				Role:    "assistant",
				Content: response,
			})
			return response
		} else {

			// 未能识别意图 依然是无意图场景回复
			response := cs.generateNoSceneResponse(ctx, userInput)
			cs.chatHistory = append(cs.chatHistory, models.ChatMessage{
				Role:    "assistant",
				Content: response,
			})
			return response
		}
	}

	// 多轮 且保持了意图场景不变
	if _, exists := cs.sceneTemplates[cs.currentPurpose]; exists {
		processor, err := cs.getProcessorForScene(ctx, cs.currentPurpose)
		if err != nil {
			return "抱歉，无法找到场景处理器。"
		}

		response, err := processor.Process(ctx, userInput, cs.chatHistory)
		if err != nil {
			response = fmt.Sprintf("处理出错: %v", err)
		}

		// 当前意图处理完成 清空场景
		if !strings.HasPrefix(response, "请问") && !strings.HasPrefix(response, "抱歉，无法找到场景") {
			cs.clearCurrentScene(ctx)
		}

		cs.chatHistory = append(cs.chatHistory, models.ChatMessage{
			Role:    "assistant",
			Content: response,
		})
		return response
	}

	// 意图不存在则构建无意图回复
	response := cs.generateNoSceneResponse(ctx, userInput)
	cs.chatHistory = append(cs.chatHistory, models.ChatMessage{
		Role:    "assistant",
		Content: response,
	})
	return response
}

func (cs *ChatbotService) recognizeIntent(ctx context.Context, userInput string) {
	purposeOptions := make(map[string]string)
	purposeDescription := make(map[string]string)
	index := 1

	// 加载意图分配  拼接Prompt  意图选项
	for _, templateInfo := range cs.sceneTemplates {
		key := fmt.Sprintf("%d", index)
		purposeOptions[key] = templateInfo.SceneName
		purposeDescription[key] = templateInfo.Description
		index++
	}

	optionsPrompt := ""
	for key, value := range purposeDescription {
		optionsPrompt += fmt.Sprintf("%s. %s - 请回复%s\n", key, value, key)
	}
	optionsPrompt += "0. 无场景/无法判断/没有符合的选项 - 请回复0"

	lastSceneInfo := "上次识别到的场景：无"
	if cs.lastRecognizedScene != "" {
		lastSceneInfo = fmt.Sprintf("上次识别到的场景：%s", cs.lastRecognizedScene)
	}

	prompt := fmt.Sprintf("有下面多种场景，需要你根据用户输入进行判断，以最新的聊天记录为准，只答选项\n%s\n用户输入：%s\n请回复序号：",
		lastSceneInfo+"\n"+optionsPrompt, userInput)

	// 拼接对话历史
	chatHistoryForLLM := make([]models.ChatMessage, len(cs.chatHistory))
	for i, msg := range cs.chatHistory {
		chatHistoryForLLM[i] = models.ChatMessage{Role: msg.Role, Content: msg.Content}
	}

	userChoice, err := llm.SendMessage(ctx, llm.SendMessageParams{
		Message:     prompt,
		UserInput:   userInput,
		ChatHistory: chatHistoryForLLM,
	})
	if err != nil {
		g.Log().Errorf(ctx, "识别意图出错: %v", err)
		return
	}

	g.Log().Printf(ctx, "purpose_options: %v", purposeOptions)
	g.Log().Printf(ctx, "user_choice: %s", userChoice)

	userChoices := utils.ExtractContinuousDigits(userChoice)

	if len(userChoices) > 0 && userChoices[0] != "0" {
		newPurpose := purposeOptions[userChoices[0]]
		if newPurpose != cs.currentPurpose {
			cs.currentPurpose = newPurpose
			cs.lastRecognizedScene = newPurpose
			cs.isSlotFilling = false
			if _, exists := cs.processors[newPurpose]; exists {
				delete(cs.processors, newPurpose)
			}
		}
		g.Log().Printf(ctx, "用户选择了场景：%s", cs.sceneTemplates[cs.currentPurpose].Name)
	} else {
		if cs.currentPurpose != "" && cs.isSlotFilling {
			g.Log().Printf(ctx, "无法判断意图，保留当前场景：%s", cs.sceneTemplates[cs.currentPurpose].Name)
		} else {
			cs.currentPurpose = ""
			cs.isSlotFilling = false
			g.Log().Printf(ctx, "无法识别用户意图\n")
		}
	}
}

// 获取场景对应的处理器
func (cs *ChatbotService) getProcessorForScene(ctx context.Context, sceneName string) (SceneProcessor, error) {

	// 根据意图获取意图对应的场景处理器
	if processor, exists := cs.processors[sceneName]; exists {
		return processor, nil
	}

	// 未找到意图对应的场景处理器 则根据场景配置 创建一个新的场景处理器
	sceneConfig, exists := cs.sceneTemplates[sceneName]
	if !exists {
		return nil, fmt.Errorf("未找到名为%s的场景配置", sceneName)
	}

	if _, exists := cs.sceneSlots[sceneName]; !exists {
		cs.sceneSlots[sceneName] = models.GetRawSlot(sceneConfig.Parameters)
	}

	processor := NewCommonProcessor(sceneConfig)
	processor.slot = cs.sceneSlots[sceneName]

	cs.processors[sceneName] = processor
	return processor, nil
}

func (cs *ChatbotService) clearCurrentScene(ctx context.Context) {
	g.Log().Printf(ctx, "场景处理，清除当前场景 %s", cs.currentPurpose)
	cs.currentPurpose = ""
	cs.lastRecognizedScene = ""
	cs.isSlotFilling = false
	cs.processors = make(map[string]SceneProcessor)
}

func (cs *ChatbotService) generateNoSceneResponse(ctx context.Context, userInput string) string {
	purposeDescription := make(map[string]string)
	index := 1

	for _, templateInfo := range cs.sceneTemplates {
		key := fmt.Sprintf("%d", index)
		purposeDescription[key] = templateInfo.Description
		index++
	}

	optionsPrompt := ""
	for key, value := range purposeDescription {
		optionsPrompt += fmt.Sprintf("%s. %s\n", key, value)
	}
	optionsPrompt += "0. 无场景/无法判断"

	prompt := prompts.GetNoSceneResponsePrompt(userInput, optionsPrompt)

	chatHistoryForLLM := make([]models.ChatMessage, len(cs.chatHistory))
	for i, msg := range cs.chatHistory {
		chatHistoryForLLM[i] = models.ChatMessage{Role: msg.Role, Content: msg.Content}
	}

	response, _ := llm.SendMessage(ctx, llm.SendMessageParams{
		Message:     prompt,
		UserInput:   userInput,
		ChatHistory: chatHistoryForLLM,
	})

	if response == "" {
		return "抱歉，我无法理解您的需求。"
	}
	return response
}

// 是否转换意图场景检查
func (cs *ChatbotService) detectSceneSwitch(ctx context.Context, userInput string) bool {
	if cs.currentPurpose == "" {
		return false
	}

	currentSceneName := cs.sceneTemplates[cs.currentPurpose].Name

	prompt := prompts.GetSceneSwitchDetectionPrompt(currentSceneName, userInput)

	chatHistoryForLLM := make([]models.ChatMessage, len(cs.chatHistory))
	for i, msg := range cs.chatHistory {
		chatHistoryForLLM[i] = models.ChatMessage{Role: msg.Role, Content: msg.Content}
	}

	response, _ := llm.SendMessage(ctx, llm.SendMessageParams{
		Message:     prompt,
		UserInput:   userInput,
		ChatHistory: chatHistoryForLLM,
	})

	digits := utils.ExtractContinuousDigits(response)
	if len(digits) > 0 && digits[0] == "1" {
		log.Printf("检测到用户意图切换场景，当前场景：%s", currentSceneName)
		return true
	}

	return false
}
