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

// 示例：在 Chain 中使用 ChatTemplate —— prompt 组件最常见的编排方式。
//
// 场景一：Chain 的输入直接就是模板变量 map，Template 后接 ChatModel。
// 场景二：上游多个节点并行产出变量，通过 outputKey 自动聚合成 map 传给 Template。
//
// 本例用 fakeModel 模拟 ChatModel；真实场景中可替换为
// eino-ext 中的 openai、ark 等实现。
//
// 运行方式：
//
//	go run ./examples/prompt/in_chain
package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
)

// fakeModel 模拟一个 ChatModel：只回显收到的最后一条消息。
type fakeModel struct{}

func (m *fakeModel) Generate(_ context.Context, input []*schema.Message,
	_ ...model.Option) (*schema.Message, error) {
	last := input[len(input)-1]
	return schema.AssistantMessage(fmt.Sprintf("(模拟回复) %s", last.Content), nil), nil
}

// Stream 是 BaseChatModel 接口的另一半，这里简单复用 Generate 实现
func (m *fakeModel) Stream(ctx context.Context, input []*schema.Message,
	opts ...model.Option) (*schema.StreamReader[*schema.Message], error) {
	msg, err := m.Generate(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	sr, sw := schema.Pipe[*schema.Message](1)
	sw.Send(msg, nil)
	sw.Close()
	return sr, nil
}

func main() {
	ctx := context.Background()

	// ------- 场景一：Chain 输入即模板变量 -------

	// ChatTemplate 在 Graph/Chain 中通常位于 ChatModel 之前
	tpl := prompt.FromMessages(schema.FString,
		schema.SystemMessage("你是{role}，请用{language}回答。"),
		schema.UserMessage("{query}"),
	)

	chain := compose.NewChain[map[string]any, *schema.Message]()
	chain.AppendChatTemplate(tpl).AppendChatModel(&fakeModel{})

	r, err := chain.Compile(ctx)
	if err != nil {
		panic(err)
	}

	out, err := r.Invoke(ctx, map[string]any{
		"role":     "eino 专家",
		"language": "中文",
		"query":    "eino 是什么？",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("场景一输出:", out.Content)

	// ------- 场景二：并行节点 + outputKey 聚合模板变量 -------

	// 模板变量 target_lang 和 query 由上游并行节点分别产出
	tpl2 := prompt.FromMessages(schema.FString,
		schema.SystemMessage("你是翻译助手，目标语言：{target_lang}。"),
		schema.UserMessage("请翻译这句话：{query}"),
	)

	parallel := compose.NewParallel().
		// AddLambda 的第一个参数是 outputKey：节点输出会以该 key 写入下游的变量 map
		AddLambda("target_lang", compose.InvokableLambda(
			func(_ context.Context, _ string) (string, error) {
				return "英文", nil
			})).
		AddLambda("query", compose.InvokableLambda(
			func(_ context.Context, in string) (string, error) {
				return in, nil
			}))

	chain2 := compose.NewChain[string, *schema.Message]()
	chain2.AppendParallel(parallel).AppendChatTemplate(tpl2).AppendChatModel(&fakeModel{})

	r2, err := chain2.Compile(ctx)
	if err != nil {
		panic(err)
	}

	out2, err := r2.Invoke(ctx, "eino 是一个用于开发 LLM 应用的 Go 框架。")
	if err != nil {
		panic(err)
	}
	fmt.Println("场景二输出:", out2.Content)
}
