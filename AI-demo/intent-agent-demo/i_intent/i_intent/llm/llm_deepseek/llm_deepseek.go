package llm_deepseek

import (
	"context"

	"github.com/cloudwego/eino-ext/components/model/deepseek"
	"github.com/cloudwego/eino/schema"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	key       string
	modelName string
	baseURL   string
)

func init() {
	ctx := context.Background()
	key = g.Cfg().MustGet(ctx, "llm.deepseek.OPENAI_API_KEY").String()
	modelName = g.Cfg().MustGet(ctx, "llm.deepseek.OPENAI_MODEL_NAME").String()
	baseURL = g.Cfg().MustGet(ctx, "llm.deepseek.OPENAI_BASE_URL").String()
}

func Chat(ctx context.Context, prompt string) (string, error) {
	chatModel, err := deepseek.NewChatModel(ctx, &deepseek.ChatModelConfig{
		Model:   modelName,
		APIKey:  key,
		BaseURL: baseURL,
	})
	if err != nil {
		return "", err
	}

	resp, err := chatModel.Generate(ctx, []*schema.Message{
		{
			Role:    schema.User,
			Content: prompt,
		},
	})
	if err != nil {
		return "", err
	}

	g.Log().Debug(ctx, "\n===================================\n")
	g.Log().Debugf(ctx, "输入的Prompt: $%s", prompt)
	g.Log().Debugf(ctx, "Deepseek resp: %v", resp)
	g.Log().Debug(ctx, "\n===================================\n")

	return resp.Content, nil
}
