package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

// getCurrentWeather 模拟获取天气信息的函数
func getCurrentWeather(location string, unit string) string {
	weatherInfo := map[string]interface{}{
		"location":    location,
		"temperature": "72",
		"unit":        unit,
		"forecast":    []string{"sunny"},
	}

	if location == "Chicago" {
		weatherInfo["temperature"] = "65"
		weatherInfo["forecast"] = []string{"windy"}
	}

	jsonData, _ := json.Marshal(weatherInfo)
	return string(jsonData)
}

// 定义工具函数
func defineTools() []llms.Tool {
	return []llms.Tool{
		{
			Type: "function",
			Function: &llms.FunctionDefinition{
				Name:        "getCurrentWeather",
				Description: "获取指定位置的当前天气信息",
				Parameters: map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"location": map[string]interface{}{
							"type":        "string",
							"description": "城市名，例如：San Francisco, CA",
						},
						"unit": map[string]interface{}{
							"type": "string",
							"enum": []string{"celsius", "fahrenheit"},
						},
					},
					"required": []string{"location"},
				},
			},
		},
	}
}

func main() {
	fmt.Println("Langchain-Go Function Call 示例")
	fmt.Println("1. 运行完整示例")
	fmt.Println("2. 运行简单示例")

	// 由于这是一个演示，我们直接运行完整示例

	// 检查是否设置了 OpenAI API 密钥
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("警告: 未设置 OPENAI_API_KEY 环境变量，将只显示代码结构")
		showCodeStructure()
		return
	}

	// 初始化 OpenAI 模型
	llm, err := openai.New(
		openai.WithModel("gpt-3.5-turbo"),
		openai.WithToken(apiKey),
	)
	if err != nil {
		log.Fatal(err)
	}

	// 定义工具
	tools := defineTools()

	// 创建对话上下文
	messages := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeHuman, "What's the weather like in Boston and Chicago?"),
	}

	ctx := context.Background()

	// 第一次调用模型，可能会触发函数调用
	completion, err := llm.GenerateContent(ctx, messages, llms.WithTools(tools))
	if err != nil {
		log.Fatal(err)
	}

	// 处理响应
	var toolCalls []llms.ToolCall
	for _, choice := range completion.Choices {
		toolCalls = append(toolCalls, choice.ToolCalls...)
		fmt.Println("AI Response:", choice.Content)
	}

	// 如果有工具调用，则执行工具
	if len(toolCalls) > 0 {
		// 将原始消息添加到历史记录中
		messages = append(messages, llms.TextParts(llms.ChatMessageTypeAI, completion.Choices[0].Content))

		// 执行所有工具调用
		for _, toolCall := range toolCalls {
			if toolCall.FunctionCall.Name == "getCurrentWeather" {
				// 解析参数
				args := toolCall.FunctionCall.Arguments
				var argsMap map[string]interface{}
				json.Unmarshal([]byte(args), &argsMap)

				location, _ := argsMap["location"].(string)
				unit, ok := argsMap["unit"].(string)
				if !ok {
					unit = "fahrenheit"
				}

				// 调用函数获取结果
				weatherResult := getCurrentWeather(location, unit)

				// 添加工具响应到消息历史
				toolResponse := llms.MessageContent{
					Role: llms.ChatMessageTypeTool,
					Parts: []llms.ContentPart{
						llms.ToolCallResponse{
							ToolCallID: toolCall.ID,
							Name:       toolCall.FunctionCall.Name,
							Content:    weatherResult,
						},
					},
				}
				messages = append(messages, toolResponse)
			}
		}

		// 再次调用模型，这次会包含工具的结果
		finalCompletion, err := llm.GenerateContent(ctx, messages, llms.WithTools(tools))
		if err != nil {
			log.Fatal(err)
		}

		// 输出最终结果
		for _, choice := range finalCompletion.Choices {
			fmt.Println("Final Response:")
			fmt.Println(choice.Content)
		}
	}
}

// showCodeStructure 展示代码结构而不实际调用 API
func showCodeStructure() {
	fmt.Println("\n=== Function Call 示例代码结构 ===")
	fmt.Println("1. 定义工具函数:")
	fmt.Println(`   - getCurrentWeather: 获取天气信息`)
	fmt.Println(`   - 参数: location (string), unit (string)`)

	fmt.Println("\n2. 创建 OpenAI 模型实例:")
	fmt.Println(`   llm, err := openai.New(openai.WithModel("gpt-3.5-turbo"))`)

	fmt.Println("\n3. 定义对话上下文:")
	fmt.Println(`   messages := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeHuman, "What's the weather like in Boston and Chicago?"),
	}`)

	fmt.Println("\n4. 首次调用模型:")
	fmt.Println(`   completion, err := llm.GenerateContent(ctx, messages, llms.WithTools(tools))`)

	fmt.Println("\n5. 处理工具调用:")
	fmt.Println(`   - 检查是否有 ToolCalls`)
	fmt.Println(`   - 解析参数并执行本地函数`)
	fmt.Println(`   - 将结果作为 ToolMessage 添加到对话历史`)

	fmt.Println("\n6. 再次调用模型:")
	fmt.Println(`   - 包含工具调用结果的完整对话历史`)
	fmt.Println(`   - 获取最终的自然语言响应`)

	fmt.Println("\n要运行此示例，请设置 OPENAI_API_KEY 环境变量。")
}
