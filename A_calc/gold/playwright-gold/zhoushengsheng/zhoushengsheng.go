package zhoushengsheng

import (
	"context"
	"github.com/playwright-community/playwright-go"
	"gold/playwright-gold/common"
	"log/slog"
)

func GetZhouShSh(ctx context.Context, page playwright.Page) error {
	addr := "https://cn.chowsangsang.com/gold-info"

	// 访问登录页面 - 将下面的URL替换为实际的登录页面
	slog.DebugContext(ctx, "正在访问登录页面 addr: %s ...", addr)
	if _, err := page.Goto(addr, playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateLoad,
	}); err != nil {
		slog.ErrorContext(ctx, "无法访问登录页面: %v", err)
		return err
	}

	slog.DebugContext(ctx, "关闭Cookie弹窗...")
	// 关闭Cookie弹窗
	cookieDialogSelector := "#onetrust-accept-btn-handler"
	_ = common.WaitForSelectorAndClick(ctx, page, cookieDialogSelector)

	slog.DebugContext(ctx, "关闭声明弹窗...")
	// 关闭声明弹窗
	statementDialogSelector := "span.close[onclick='closeEmergencyPopup()']"
	_ = common.WaitForSelectorAndClick(ctx, page, statementDialogSelector)

	// 等待表格加载完成
	err := common.WaitForSelector(page, ".table-responsive.table-size")
	if err != nil {
		slog.ErrorContext(ctx, "等待元素[%s]失败: %v", "table of gold", err)
		return err
	}

	return nil
}
