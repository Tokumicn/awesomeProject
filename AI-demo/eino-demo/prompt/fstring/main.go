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

// 示例：FString 模板 —— 最常用的模板格式。
//
// FString 使用 Python 风格的 {var} 占位符，通过 prompt.FromMessages 创建模板，
// 调用 Format 用变量渲染出发送给 ChatModel 的消息列表。
//
// 运行方式：
//
//	go run ./examples/prompt/fstring
package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
)

func main() {
	ctx := context.Background()

	// 1. 构造模板：格式类型（schema.FString）+ 一组消息模板。
	//    消息顺序就是渲染后的消息顺序，通常为 System 在前、User 在后。
	tpl := prompt.FromMessages(schema.FString,
		schema.SystemMessage("你是{role}，请用{language}回答。"),
		schema.UserMessage("{query}"),
	)

	// 2. 渲染：传入变量 map，占位符会被替换成对应的值。
	msgs, err := tpl.Format(ctx, map[string]any{
		"role":     "eino 专家",
		"language": "中文",
		"query":    "eino 是什么？",
	})
	if err != nil {
		panic(err)
	}

	// 3. 渲染结果即为可直接传给 ChatModel.Generate 的消息列表。
	for _, msg := range msgs {
		fmt.Printf("[%s] %s\n", msg.Role, msg.Content)
	}

	// 4. 注意：模板中出现的变量在渲染时必须全部提供，缺失会返回错误。
	//    （下面只提供了 query，缺少 role 和 language）
	_, err = tpl.Format(ctx, map[string]any{
		"query": "你好",
	})
	fmt.Printf("\n缺少变量时报错: %v\n", err)
}
