package main

import (
	"context"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
)

// 获取鼓励师提示模板
func getEncouragementPrompt() ([]*schema.Message, error) {
	// 创建模板，使用 FString 格式
	template := prompt.FromMessages(schema.FString,
		// 系统消息模板
		schema.SystemMessage("你是一个{role}。你需要用{style}的语气回答问题。你的目标是帮助程序员保持积极乐观的心态，提供技术建议的同时也要关注他们的心理健康。"),

		// 插入需要的对话历史（新对话的话这里不填）
		schema.MessagesPlaceholder("chat_history", true),

		// 用户消息模板
		schema.UserMessage("问题: {question}"),
	)

	// 使用模板生成消息
	messages, err := template.Format(context.Background(), map[string]any{
		"role":     "程序员鼓励师",
		"style":    "积极、温暖且专业",
		"question": "我的代码一直报错，感觉好沮丧，该怎么办？",
		// 对话历史（这个例子里模拟两轮对话历史）
		"chat_history": []*schema.Message{
			schema.UserMessage("你好"),
			schema.AssistantMessage("嘿！我是你的程序员鼓励师！记住，每个优秀的程序员都是从 Debug 中成长起来的。有什么我可以帮你的吗？", nil),
			schema.UserMessage("我觉得自己写的代码太烂了"),
			schema.AssistantMessage("每个程序员都经历过这个阶段！重要的是你在不断学习和进步。让我们一起看看代码，我相信通过重构和优化，它会变得更好。记住，Rome wasn't built in a day，代码质量是通过持续改进来提升的。", nil),
		},
	})
	if err != nil {
		return nil, err
	}

	return messages, nil
}

type ChatHistory struct {
	UserA   string `json:"usr_a"`
	SystemQ string `json:"system_q"`
}

// 对话重写
func getQueryRewritePrompt(query string, histories []ChatHistory) ([]*schema.Message, error) {
	// 创建模板，使用 FString 格式
	template := prompt.FromMessages(schema.FString,
		// 系统消息模板
		schema.SystemMessage("你是一个{role}。根据历史对话记录，对用户问题进行补全，尽量有限使用最近的对话内容进行改写，不要任意发挥。"),

		// 插入需要的对话历史（新对话的话这里不填）
		schema.MessagesPlaceholder("chat_history", true),

		// 用户消息模板
		schema.UserMessage("问题: {question}"),
	)

	// 提供历史信息
	his := []*schema.Message{}
	for _, h := range histories {
		his = append(his, schema.UserMessage(h.UserA))
		his = append(his, schema.UserMessage(h.SystemQ))
	}

	// 使用模板生成消息
	messages, err := template.Format(context.Background(), map[string]any{
		"role":     "用户问题重写人员",
		"question": query,
		// 对话历史（这个例子里模拟两轮对话历史）
		//"chat_history": []*schema.Message{
		//	schema.UserMessage("第一名奖品是什么？"),
		//	schema.AssistantMessage("第一名奖品是第一名", nil),
		//	schema.UserMessage("第二名呢？"),
		//	schema.AssistantMessage("奖品是第二名", nil),
		//},
		"chat_history": his,
	})
	if err != nil {
		return nil, err
	}

	return messages, nil
}
