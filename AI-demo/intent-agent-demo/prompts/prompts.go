package prompts

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"

	"text/template"

	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
)

func BasePrompt() ([]*schema.Message, error) {

	systemMsg := "你是一个{role}。"
	placeholderKey := "history_key"
	taskMsg := "请帮我{task}。"
	// 准备变量
	variables := map[string]any{
		"role": "专业的助手",
		"task": "写一首诗",
		"history_key": []*schema.Message{
			{Role: schema.User, Content: "告诉我油画是什么?"},
			{Role: schema.Assistant, Content: "油画是xxx"},
		},
	}

	messages, err := basePrompt(systemMsg, placeholderKey, taskMsg, variables)
	if err != nil {
		return nil, err
	}

	// 打印格式化后的消息
	for _, msg := range messages {
		fmt.Println("=================")
		fmt.Println("[DEBUG] ", msg.Content)
		fmt.Println("=================")
	}

	return messages, nil
}

func basePrompt(systemMsg, placeholderKey, task string, variables map[string]any) ([]*schema.Message, error) {
	// 创建模板
	template := prompt.FromMessages(schema.FString,
		schema.SystemMessage(systemMsg),
		schema.MessagesPlaceholder(placeholderKey, false),
		&schema.Message{
			Role:    schema.User,
			Content: task,
		},
	)

	// 格式化模板
	messages, err := template.Format(context.Background(), variables)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func GoTemplateFormat(tmpl string, variables map[string]any) (string, error) {

	type PromptData struct {
		Query   string   `json:"query,omitempty"`
		History []string `json:"history,omitempty"`
	}

	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	filePath := filepath.Join(dir, "/gotemplate_demo.tmpl")

	goPromptStr, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	tmpl = string(goPromptStr)

	data := PromptData{
		Query: "怎么买比特比？",
		History: []string{
			"用户之前询问过数字货币相关的问题。",
		},
	}

	t := template.Must(template.New("prompt").Parse(tmpl))

	var buf bytes.Buffer
	err = t.Execute(&buf, data)
	if err != nil {
		panic(err)
	}

	renderedPrompt := buf.String()
	fmt.Println(renderedPrompt)
	return renderedPrompt, nil
}

func Jinja2Format(template string, variables map[string]any) (string, error) {

	return "", nil
}
