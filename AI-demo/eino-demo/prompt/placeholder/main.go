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

// 示例：MessagesPlaceholder —— 在模板中插入一组完整消息（如多轮对话历史）。
//
// 普通消息模板只能做"字符串替换"，而 MessagesPlaceholder 可以把
// 变量中的 []*schema.Message 原样插入消息列表的对应位置，
// 常用于拼接历史对话（history / few-shot examples）。
//
// 运行方式：
//
//	go run ./examples/prompt/placeholder
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
		schema.SystemMessage("你是 eino 助手，请结合历史对话回答。"),

		// 必填占位符：变量中必须存在 key 为 "history" 的 []*schema.Message，
		// 缺失时 Format 返回错误。
		schema.MessagesPlaceholder("history", false),

		schema.UserMessage("{query}"),

		// 可选占位符：变量缺失时渲染为空，不报错（optional = true）。
		schema.MessagesPlaceholder("examples", true),
	)

	msgs, err := tpl.Format(ctx, map[string]any{
		"history": []*schema.Message{
			schema.UserMessage("eino 是什么？"),
			schema.AssistantMessage("eino 是一个用于开发 LLM 应用的 Go 框架。", nil),
		},
		"query": "它支持哪些组件？",
		// 故意不提供 "examples"，因为是可选占位符
	})
	if err != nil {
		panic(err)
	}

	for _, msg := range msgs {
		fmt.Printf("[%s] %s\n", msg.Role, msg.Content)
	}
}
