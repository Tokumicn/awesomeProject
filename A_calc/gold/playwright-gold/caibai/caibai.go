package caibai

import (
	"context"
	"github.com/playwright-community/playwright-go"
	"log/slog"
)

func GetCaiBai(ctx context.Context, page playwright.Page) error {
	addr := "http://cbwx.bjcaibai.com.cn/wbap/#/"

	// 访问登录页面 - 将下面的URL替换为实际的登录页面
	slog.DebugContext(ctx, "正在访问登录页面 addr: %s ...", addr)
	if _, err := page.Goto(addr, playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateNetworkidle,
	}); err != nil {
		slog.ErrorContext(ctx, "无法访问登录页面: %v", err)
		return err
	}

	return nil
}
