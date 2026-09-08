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

// 示例：GoTemplate 模板 —— 需要控制流（if / range）时选用。
//
// GoTemplate 使用 Go 标准库 text/template 语法（{{.var}}），
// 支持 range、if、len 等能力，适合在模板里做循环和条件拼接。
// 注意：变量缺失时同样会报错（missingkey=error）。
//
// 运行方式：
//
//	go run ./examples/prompt/go_template
package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
)

func main() {
	ctx := context.Background()

	tpl := prompt.FromMessages(schema.GoTemplate,
		schema.SystemMessage("你是翻译助手，目标语言：{{.language}}。"),
		schema.UserMessage("请翻译以下 {{len .words}} 个单词：{{range .words}}「{{.}}」{{end}}"),
	)

	msgs, err := tpl.Format(ctx, map[string]any{
		"language": "英文",
		"words":    []string{"苹果", "香蕉", "橘子"},
	})
	if err != nil {
		panic(err)
	}

	for _, msg := range msgs {
		fmt.Printf("[%s] %s\n", msg.Role, msg.Content)
	}
}
