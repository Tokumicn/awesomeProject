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

// 示例：Agentic 消息模板 —— 面向 Agentic 编排的消息模板。
//
// prompt.FromAgenticMessages 与 FromMessages 用法一致，
// 区别在于渲染结果是 []*schema.AgenticMessage（带内容块的 agentic 消息），
// 可配合 AgenticModel / AgenticChatTemplate 节点在 Graph 中使用。
//
// 运行方式：
//
//	go run ./examples/prompt/agentic
package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
)

func main() {
	ctx := context.Background()

	tpl := prompt.FromAgenticMessages(schema.FString,
		schema.SystemAgenticMessage("你是 eino 助手，请简洁回答。"),

		// 与 MessagesPlaceholder 对应，这里渲染 []*schema.AgenticMessage
		schema.AgenticMessagesPlaceholder("history", false),

		schema.UserAgenticMessage("{query}"),
	)

	msgs, err := tpl.Format(ctx, map[string]any{
		"history": []*schema.AgenticMessage{
			schema.UserAgenticMessage("eino 是什么？"),
			// assistant 消息需要手动构造（目前只有 System/User 构造器）
			{
				Role: schema.AgenticRoleTypeAssistant,
				ContentBlocks: []*schema.ContentBlock{
					schema.NewContentBlock(&schema.UserInputText{
						Text: "eino 是一个用于开发 LLM 应用的 Go 框架。",
					}),
				},
			},
		},
		"query": "它支持哪些组件？",
	})
	if err != nil {
		panic(err)
	}

	for _, msg := range msgs {
		fmt.Printf("[%s]\n", msg.Role)
		for _, block := range msg.ContentBlocks {
			if block.UserInputText != nil {
				fmt.Printf("  text: %s\n", block.UserInputText.Text)
			}
		}
	}
}
