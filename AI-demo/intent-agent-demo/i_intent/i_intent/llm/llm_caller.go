package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"i_intent/data/model"

	"github.com/gogf/gf/v2/frame/g"
)

type LLMRequest struct {
	Model          string               `json:"model"`
	Messages       []models.ChatMessage `json:"messages"`
	EnableThinking bool                 `json:"enable_thinking"`
}

type LLMResponse struct {
	Choices []struct {
		Message models.ChatMessage `json:"message"`
	} `json:"choices"`
}

type SendMessageParams struct {
	Message     string
	UserInput   string
	ChatHistory []models.ChatMessage
}

func SendMessage(ctx context.Context, params SendMessageParams) (string, error) {
	fmt.Println("--------------------------------------------------------------------")
	fmt.Println("prompt输入:", params.Message)
	fmt.Println("----------------------------------")

	systemPrompt := GetSystemPrompt()
	modelName := GetModel(ctx)
	url := GetGPTUrl(ctx)
	apiKey := GetAPIKey(ctx)
	timeOut := GetAPITimeout(ctx)

	messages := []models.ChatMessage{
		{Role: "system", Content: systemPrompt},
	}

	if params.ChatHistory != nil {
		for _, msg := range params.ChatHistory {
			messages = append(messages, models.ChatMessage{Role: msg.Role, Content: msg.Content})
		}
	}

	messages = append(messages, models.ChatMessage{Role: "user", Content: params.Message})

	req := LLMRequest{
		Model:          modelName,
		Messages:       messages,
		EnableThinking: false,
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}

	httpReq.Header.Set("Authorization", "Bearer "+apiKey)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: time.Duration(timeOut) * time.Second}
	resp, err := client.Do(httpReq)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("API error: status %d, body: %s", resp.StatusCode, string(body))
	}

	var llmResp LLMResponse
	if err := json.Unmarshal(body, &llmResp); err != nil {
		return "", err
	}

	if len(llmResp.Choices) == 0 {
		return "", fmt.Errorf("no choices in response")
	}

	answer := llmResp.Choices[0].Message.Content
	fmt.Println("LLM输出:", answer)
	fmt.Println("--------------------------------------------------------------------")

	return answer, nil
}

func CallSceneAPI(ctx context.Context, sceneName string, slotsData map[string]interface{}) (map[string]interface{}, error) {
	apiURL := fmt.Sprintf(GetSceneAPIURLTemplate(), sceneName)

	reqBody, err := json.Marshal(slotsData)
	if err != nil {
		return nil, err
	}

	fmt.Printf("调用场景API: %s\n", apiURL)
	fmt.Printf("请求体: %s\n", string(reqBody))

	httpReq, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: time.Duration(GetAPITimeout(ctx)) * time.Second}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return map[string]interface{}{"error": fmt.Sprintf("API调用失败，状态码: %d", resp.StatusCode)}, nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	fmt.Printf("API调用成功: %s\n", string(body))
	return result, nil
}

func ProcessAPIResult(ctx context.Context, apiResult map[string]interface{}, chatHistory []models.ChatMessage) (string, error) {
	dataPart, ok := apiResult["data"]
	if !ok {
		dataPart = apiResult
	}

	dataJSON, err := json.Marshal(dataPart)
	if err != nil {
		return "", err
	}

	prompt := fmt.Sprintf(GetAPIResultPrompt(), string(dataJSON))

	result, err := SendMessage(ctx, SendMessageParams{
		Message:     prompt,
		UserInput:   "",
		ChatHistory: chatHistory,
	})

	if result == "" {
		return "抱歉，处理结果时出现错误，请稍后重试。", nil
	}

	return result, err
}

func GetSystemPrompt() string { return "You are a helpful assistant." }

func GetModel(ctx context.Context) string {
	return g.Cfg().MustGet(ctx, "llm.deepseek.OPENAI_MODEL_NAME").String()
}

func GetGPTUrl(ctx context.Context) string {
	return g.Cfg().MustGet(ctx, "llm.deepseek.OPENAI_BASE_URL").String()
}

func GetAPIKey(ctx context.Context) string {
	return g.Cfg().MustGet(ctx, "llm.deepseek.OPENAI_API_KEY").String()
}

func GetAPITimeout(ctx context.Context) int {
	return g.Cfg().MustGet(ctx, "llm.deepseek.OPENAI_API_TIMEOUT").Int()
}
func GetSceneAPIURLTemplate() string { return "https://example.com/api/{}" }

func GetAPIResultPrompt() string {
	return "根据以下API返回结果生成用户友好的回复：\n{}"
}
