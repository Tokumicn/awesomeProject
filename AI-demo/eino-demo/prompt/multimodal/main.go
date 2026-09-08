/*
 * Copyright 2026 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// 示例：多模态模板 —— 渲染带图片等多模态内容的消息。
//
// 使用 schema.Message 的 MultiContent 字段可以定义图文混合的消息模板，
// 文本与图片 URL 中的占位符都会被渲染。渲染结果可直接传给
// 支持视觉的 ChatModel（如 gpt-4o、doubao-vision 等）。
//
// 运行方式：
//
//	go run ./examples/prompt/multimodal
package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
)

func main() {
	ctx := context.Background()

	tpl := prompt.FromMessages(schema.FString,
		// 多模态消息模板：文本 part 和图片 part 都支持占位符
		&schema.Message{
			Role: schema.User,
			MultiContent: []schema.ChatMessagePart{
				{
					Type: schema.ChatMessagePartTypeText,
					Text: "请{action}这张图片",
				},
				{
					Type:     schema.ChatMessagePartTypeImageURL,
					ImageURL: &schema.ChatMessageImageURL{URL: "{image_url}"},
				},
			},
		},
	)

	msgs, err := tpl.Format(ctx, map[string]any{
		"action":    "描述",
		"image_url": "https://example.com/cat.png",
	})
	if err != nil {
		panic(err)
	}

	for _, msg := range msgs {
		fmt.Printf("[%s] %d 个内容块:\n", msg.Role, len(msg.MultiContent))
		for _, part := range msg.MultiContent {
			switch part.Type {
			case schema.ChatMessagePartTypeText:
				fmt.Printf("  - text: %s\n", part.Text)
			case schema.ChatMessagePartTypeImageURL:
				fmt.Printf("  - image_url: %s\n", part.ImageURL.URL)
			}
		}
	}
}
