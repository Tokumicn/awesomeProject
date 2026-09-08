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

// 示例：自定义 ChatTemplate —— 实现自己的模板逻辑并支持调用时选项。
//
// 只要实现 prompt.ChatTemplate 接口（一个 Format 方法），就可以：
//   - 在 Graph/Chain 中与内置模板一样通过 AppendChatTemplate 编排
//   - 通过 prompt.WrapImplSpecificOptFn / GetImplSpecificOptions
//     定义"实现自定义"的调用时 Option
//
// 运行方式：
//
//	go run ./examples/prompt/custom_template
package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
)

// myOptions 是本实现自定义的调用时选项，带有默认值。
type myOptions struct {
	Uppercase bool // 渲染后是否转大写
	Prefix    string
}

// myTemplate 实现了 prompt.ChatTemplate 接口。
type myTemplate struct{}

// Format 把变量渲染成消息列表。
// 通过 prompt.GetImplSpecificOptions 提取实现自定义的选项，并应用默认值。
func (t *myTemplate) Format(_ context.Context, vs map[string]any,
	opts ...prompt.Option) ([]*schema.Message, error) {
	// base 中的值是默认值，调用方传入的 opts 会覆盖它
	o := prompt.GetImplSpecificOptions(&myOptions{Prefix: "用户提问"}, opts...)

	query, _ := vs["query"].(string)
	if o.Uppercase {
		query = strings.ToUpper(query)
	}

	return []*schema.Message{
		schema.SystemMessage("你是 eino 助手。"),
		schema.UserMessage(o.Prefix + ": " + query),
	}, nil
}

func main() {
	ctx := context.Background()

	tpl := &myTemplate{}

	// 1. 不传选项：使用默认值（Prefix = "用户提问"，Uppercase = false）
	msgs, err := tpl.Format(ctx, map[string]any{"query": "eino 是什么？"})
	if err != nil {
		panic(err)
	}
	fmt.Println("默认选项渲染结果:")
	for _, msg := range msgs {
		fmt.Printf("  [%s] %s\n", msg.Role, msg.Content)
	}

	// 2. 用 WrapImplSpecificOptFn 构造实现自定义的选项，在调用时覆盖默认值
	uppercaseOpt := prompt.WrapImplSpecificOptFn(func(o *myOptions) {
		o.Uppercase = true
	})
	prefixOpt := prompt.WrapImplSpecificOptFn(func(o *myOptions) {
		o.Prefix = "原始输入"
	})

	msgs, err = tpl.Format(ctx, map[string]any{"query": "eino 是什么？"}, uppercaseOpt, prefixOpt)
	if err != nil {
		panic(err)
	}
	fmt.Println("自定义选项渲染结果:")
	for _, msg := range msgs {
		fmt.Printf("  [%s] %s\n", msg.Role, msg.Content)
	}
}
