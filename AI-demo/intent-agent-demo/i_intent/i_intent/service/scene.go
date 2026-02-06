package service

import (
	"context"
	"fmt"
	"i_intent/data/model"
	"i_intent/llm"
	"i_intent/llm/prompts"
	"i_intent/utils"

	"github.com/gogf/gf/v2/frame/g"
)

var _commonCommonProcessor *CommonProcessor

type SceneProcessor interface {
	Process(ctx context.Context, userInput string, chatHistory []models.ChatMessage) (string, error)
}

type CommonProcessor struct {
	sceneConfig        models.SceneConfig
	sceneName          string
	slotTemplate       []models.Slot
	slotDynamicExample string
	slot               []models.Slot
}

func NewCommonProcessor(sceneConfig models.SceneConfig) *CommonProcessor {
	if _commonCommonProcessor == nil {
		_commonCommonProcessor = &CommonProcessor{
			sceneConfig:        sceneConfig,
			sceneName:          sceneConfig.Name,
			slotTemplate:       models.GetRawSlot(sceneConfig.Parameters),
			slotDynamicExample: models.GetDynamicExample(sceneConfig),
			slot:               models.GetRawSlot(sceneConfig.Parameters),
		}
	}
	return _commonCommonProcessor
}

func (cp *CommonProcessor) Process(ctx context.Context, userInput string, chatHistory []models.ChatMessage) (string, error) {
	slotTemplateStr := make([]string, 0, len(cp.slotTemplate))
	for _, s := range cp.slotTemplate {
		slotData := map[string]interface{}{
			"name":  s.Name,
			"desc":  s.Desc,
			"type":  s.Type,
			"value": s.Value,
		}
		jsonBytes, _ := utils.FormatSlotJSON(slotData)
		slotTemplateStr = append(slotTemplateStr, string(jsonBytes))
	}

	// 根据词槽配置 构建提取词槽Prompt
	message := prompts.GetSlotUpdateMessage(cp.sceneName, cp.slotDynamicExample, slotTemplateStr, userInput)

	llmChatHistory := make([]models.ChatMessage, len(chatHistory))
	for i, msg := range chatHistory {
		llmChatHistory[i] = models.ChatMessage{Role: msg.Role, Content: msg.Content}
	}

	newInfoJSONRaw, err := llm.SendMessage(ctx, llm.SendMessageParams{
		Message:     message,
		UserInput:   userInput,
		ChatHistory: llmChatHistory,
	})
	if err != nil {
		return "", err
	}

	currentValues := utils.ExtractJSONFromString(newInfoJSONRaw)
	if len(currentValues) > 0 {
		if _, ok := currentValues[0]["name"]; !ok {
			newValues := make([]map[string]interface{}, 0)
			newValues = append(newValues, currentValues[0])
			models.UpdateSlot(newValues, cp.slot)
		} else {
			models.UpdateSlot(currentValues, cp.slot)
		}
	}

	g.Log().Printf(ctx, "slot update before: %v", cp.slot)
	models.UpdateSlot(currentValues, cp.slot)
	g.Log().Printf(ctx, "slot update after: %v", cp.slot)

	// 检查是否必须得词槽都已填充
	if models.IsSlotFullyFilled(cp.slot) {
		// 所有必须词槽都已填充，返回完整数据
		return cp.respondWithCompleteData(ctx, chatHistory)
	}
	// 词槽缺失  追加反问
	return cp.askUserForMissingData(ctx, userInput, chatHistory)
}

func (cp *CommonProcessor) formatSlot(slot models.Slot) ([]byte, error) {
	return nil, nil
}

func (cp *CommonProcessor) respondWithCompleteData(ctx context.Context, chatHistory []models.ChatMessage) (string, error) {
	g.Log().Printf(ctx, "%s ------ 参数已完整，详细参数如下", cp.sceneName)
	g.Log().Printf(ctx, models.FormatNameValueForLogging(cp.slot))
	g.Log().Printf(ctx, "正在请求%sAPI，请稍后……", cp.sceneName)

	sceneKey := cp.getSceneKey()
	if sceneKey == "" {
		return fmt.Sprintf("抱歉，无法找到场景 '%s' 的配置信息。", cp.sceneName), nil
	}

	slotsData := make(map[string]interface{})
	for _, slot := range cp.slot {
		if slot.Value != "" {
			slotKey := cp.getSlotKey(slot.Name)
			if slotKey != "" {
				slotsData[slotKey] = slot.Value
			}
		}
	}

	// TODO 无需真的调用API 数据返回即可
	apiResult := map[string]interface{}{
		"data": fmt.Sprintf("成功调用了API 请求的场景是[%s]，参数如下：\n[%v]\n", cp.sceneName, slotsData),
	}

	//apiResult, err := llm.CallSceneAPI(ctx, sceneKey, slotsData)
	//if err != nil {
	//	return fmt.Sprintf("抱歉，调用API时出现错误：%v", err), nil
	//}

	if _, ok := apiResult["error"]; ok {
		return fmt.Sprintf("抱歉，调用API时出现错误：%v", apiResult["error"]), nil
	}

	chatHistoryForLLM := make([]models.ChatMessage, len(chatHistory))
	for i, msg := range chatHistory {
		chatHistoryForLLM[i] = models.ChatMessage{Role: msg.Role, Content: msg.Content}
	}

	// 大模型处理API结果，返回用户友好的响应
	userFriendlyResponse, err := llm.ProcessAPIResult(ctx, apiResult, chatHistoryForLLM)
	if err != nil {
		return fmt.Sprintf("抱歉，处理API结果时出现错误：%v", err), nil
	}

	return userFriendlyResponse, nil
}

func (cp *CommonProcessor) askUserForMissingData(ctx context.Context, userInput string, chatHistory []models.ChatMessage) (string, error) {
	slotTemplateStr := make([]string, 0, len(cp.slot))
	for _, s := range cp.slot {
		if s.Value == "" {
			slotData := map[string]interface{}{
				"name":  s.Name,
				"desc":  s.Desc,
				"type":  s.Type,
				"value": s.Value,
			}
			jsonBytes, _ := utils.FormatSlotJSON(slotData)
			slotTemplateStr = append(slotTemplateStr, string(jsonBytes))
		}
	}

	message := prompts.GetSlotQueryUserMessage(cp.sceneName, slotTemplateStr, userInput)

	llmChatHistory := make([]models.ChatMessage, len(chatHistory))
	for i, msg := range chatHistory {
		llmChatHistory[i] = models.ChatMessage{Role: msg.Role, Content: msg.Content}
	}

	result, err := llm.SendMessage(ctx, llm.SendMessageParams{
		Message:     message,
		UserInput:   userInput,
		ChatHistory: llmChatHistory,
	})

	return result, err
}

func (cp *CommonProcessor) getSceneKey() string {
	return cp.sceneConfig.SceneName
}

func (cp *CommonProcessor) getSlotKey(slotName string) string {
	for _, param := range cp.sceneConfig.Parameters {
		if param.Name == slotName {
			return param.Name
		}
	}
	return slotName
}
