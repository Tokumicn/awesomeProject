package laomiao

import (
	"context"
	"github.com/playwright-community/playwright-go"
	"gold/playwright-gold/common"
	"log/slog"
)

func GetLaoMiao(ctx context.Context, page playwright.Page) error {
	addr := "https://www.laomiao.com.cn/glod.html"

	// 访问登录页面 - 将下面的URL替换为实际的登录页面
	slog.DebugContext(ctx, "正在访问登录页面 addr: %s ...", addr)
	if _, err := page.Goto(addr, playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateLoad,
	}); err != nil {
		slog.ErrorContext(ctx, "无法访问登录页面: %v", err)
		return err
	}

	// 等待表格加载完成
	selector := "#gold_date" // ID选择器
	err := common.WaitForSelector(page, selector)
	if err != nil {
		slog.ErrorContext(ctx, "等待元素[%s]失败: %v", selector, err)
		return err
	}

	return nil
}
