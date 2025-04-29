package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/eino-ext/components/model/ollama"
	"github.com/cloudwego/eino/schema"
	"io"
	"log"
)

func main() {
	ctx := context.TODO()
	chatModel, err := ollama.NewChatModel(ctx, &ollama.ChatModelConfig{
		BaseURL: "http://localhost:11434", // Ollama 服务地址
		Model:   "qwen3:0.6b",             // 模型名称
	})
	if err != nil {
		log.Printf("创建模型失败: %v\n", err)
		return
	}

	// 程序员鼓励师
	// messages, err := getEncouragementPrompt()

	messages, err := getQueryRewritePrompt("")

	result, err := chatModel.Generate(ctx, messages)
	fmt.Println(result)

	// 流式结果
	//streamResult, err := chatModel.Stream(ctx, messages)
}

func reportStream(sr *schema.StreamReader[*schema.Message]) {
	defer sr.Close()

	i := 0
	for {
		message, err := sr.Recv()
		if err == io.EOF { // 流式输出结束
			return
		}
		if err != nil {
			log.Fatalf("recv failed: %v", err)
		}
		log.Printf("message[%d]: %+v\n", i, message)
		i++
	}
}
