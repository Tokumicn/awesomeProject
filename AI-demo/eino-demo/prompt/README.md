# Prompt 组件示例

[eino](https://github.com/cloudwego/eino) 中 Prompt 组件（`components/prompt`）负责把变量 map 渲染成发送给 ChatModel 的消息列表。

## 核心概念

```go
// 1. 构造模板：格式类型 + 一组消息模板
tpl := prompt.FromMessages(schema.FString,
    schema.SystemMessage("你是{role}。"),
    schema.MessagesPlaceholder("history", false), // 消息占位符
    schema.UserMessage("{query}"),
)

// 2. 渲染：传入变量，得到 []*schema.Message
msgs, err := tpl.Format(ctx, map[string]any{
    "role":    "eino 专家",
    "history": historyMsgs,
    "query":   "eino 是什么？",
})

// 3. 送入模型
resp, err := chatModel.Generate(ctx, msgs)
```

三种模板格式（`schema.FormatType`）：

| 格式 | 语法 | 适用场景 |
|---|---|---|
| `schema.FString` | `{var}`（Python format 风格） | 最常用，简单变量替换 |
| `schema.GoTemplate` | `{{.var}}`（text/template） | 需要 if / range 等控制流 |
| `schema.Jinja2` | `{{ var }}`、`{% for %}` | 表达能力最强，适合外部导入的模板 |

## 示例列表

| 示例 | 说明 |
|---|---|
| [fstring](./fstring) | FString 基础用法：构造模板、渲染变量、缺变量报错 |
| [go_template](./go_template) | GoTemplate：在模板中使用 `range` / `len` 控制流 |
| [jinja2](./jinja2) | Jinja2：使用 `{{ }}` / `{% for %}` 语法的动态模板 |
| [placeholder](./placeholder) | MessagesPlaceholder：插入多轮对话历史等消息列表（必填/可选） |
| [multimodal](./multimodal) | 多模态模板：渲染图文混合（MultiContent）的消息 |
| [in_chain](./in_chain) | 在 compose.Chain 中编排：Template → ChatModel，以及并行节点通过 outputKey 聚合模板变量 |
| [custom_template](./custom_template) | 自定义 ChatTemplate 实现，并通过 WrapImplSpecificOptFn 支持调用时选项 |
| [agentic](./agentic) | Agentic 模板：FromAgenticMessages 渲染 []*schema.AgenticMessage |

## 运行方式

在仓库根目录执行（无需任何外部依赖，模型均为本地 mock）：

```bash
go run ./examples/prompt/fstring
go run ./examples/prompt/go_template
go run ./examples/prompt/jinja2
go run ./examples/prompt/placeholder
go run ./examples/prompt/multimodal
go run ./examples/prompt/in_chain
go run ./examples/prompt/custom_template
go run ./examples/prompt/agentic
```

## 注意事项

- 模板中出现的变量在 `Format` 时必须全部提供（`MessagesPlaceholder` 设为 optional 除外），缺失会返回错误，且没有编译期检查，需保证模板与调用方的变量命名一致。
- 在 Graph/Chain 中，ChatTemplate 通常位于 ChatModel 之前；上游多个节点可通过 `WithOutputKey`（Chain 的 `Parallel.AddXxx` 已内置）把输出聚合成模板变量 map。
- 本目录示例中的 ChatModel 均为 mock 实现。接入真实模型（OpenAI、火山方舟等）参见 [eino-ext](https://github.com/cloudwego/eino-ext)。
