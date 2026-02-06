package main

import (
	"context"
	"fmt"
	"log"
	"os"

	_ "intent-agent-demo/env"

	"github.com/cloudwego/eino-ext/components/model/deepseek"
	"github.com/cloudwego/eino/schema"
)

func main() {

	key := os.Getenv("OPENAI_API_KEY")
	modelName := os.Getenv("OPENAI_MODEL_NAME")
	baseURL := os.Getenv("OPENAI_BASE_URL")

	ctx := context.Background()
	chatModel, err := deepseek.NewChatModel(ctx, &deepseek.ChatModelConfig{
		Model:   modelName,
		APIKey:  key,
		BaseURL: baseURL,
	})
	if err != nil {
		log.Fatal(err)
	}

	resp, err := chatModel.Generate(ctx, []*schema.Message{
		{
			Role:    schema.User,
			Content: "你好",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
