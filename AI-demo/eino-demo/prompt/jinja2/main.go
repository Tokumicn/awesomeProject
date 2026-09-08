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

// 示例：Jinja2 模板 —— 复杂动态提示词时选用。
//
// Jinja2 使用 Jinja 模板语法（{{ var }}、{% for %}、{% if %} 等），
// 表达能力最强，适合从外部（如 Prompt 平台）导入的模板。
//
// 运行方式：
//
//	go run ./examples/prompt/jinja2
package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
)

func main() {
	ctx := context.Background()

	tpl := prompt.FromMessages(schema.Jinja2,
		schema.SystemMessage("你是会议纪要助手，请用{{ style }}的风格输出。"),
		schema.UserMessage(`请把以下 {{ n }} 条要点整理成一段话：
{% for item in items %}
- {{ item }}
{% endfor %}`),
	)

	msgs, err := tpl.Format(ctx, map[string]any{
		"style": "简洁",
		"n":     3,
		"items": []string{
			"本周上线模板功能",
			"修复两个线上问题",
			"下周一评审新方案",
		},
	})
	if err != nil {
		panic(err)
	}

	for _, msg := range msgs {
		fmt.Printf("[%s] %s\n", msg.Role, msg.Content)
	}
}
