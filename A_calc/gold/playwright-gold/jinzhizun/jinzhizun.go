package jinzhizun

import (
	"context"
	"github.com/playwright-community/playwright-go"
	"gold/playwright-gold/common"
	"log/slog"
)

func GetJinZhiZun(ctx context.Context, page playwright.Page) error {
	addr := "https://www.3dg-group.hk/"

	// 访问登录页面 - 将下面的URL替换为实际的登录页面
	slog.DebugContext(ctx, "正在访问登录页面 addr: ", addr)
	if _, err := page.Goto(addr, playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateDomcontentloaded,
	}); err != nil {
		slog.ErrorContext(ctx, "无法访问登录页面: ", err.Error())
		return err
	}

	menuOpenSelector := ".menu_open"
	err := common.WaitForSelectorAndClick(ctx, page, menuOpenSelector)
	if err != nil {
		slog.ErrorContext(ctx, "点击今日金价失败: ", err.Error())
	}

	// 点击今日金价
	goldPriceSelector := ".list goldPrice"
	err = common.WaitForSelectorAndClick(ctx, page, goldPriceSelector)
	if err != nil {
		slog.ErrorContext(ctx, "点击今日金价失败: ", err.Error())
	}

	return nil
}
