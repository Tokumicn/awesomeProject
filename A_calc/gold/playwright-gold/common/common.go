package common

import (
	"context"
	"github.com/playwright-community/playwright-go"
	"log/slog"
)

// 选择器说明
// selector := "#labTitle" // id选择器
// selector := ".gold_price_time" // class选择器

const (
	ShowChrome = true             // 是否显示浏览器
	LaoFX      = "laofengxiang"   // 老凤祥
	ZhouSS     = "zhoushengsheng" // 周生生
	ZhouLF     = "zhouliufu"      // 周六福
	ZhouDF     = "zhoudafu"       // 周大福
	LaoM       = "laomiao"        // 老庙
	CaiBai     = "caibai"         // 菜百黄金
	JinZZ      = "jinzhizun"      // 金至尊
)

func WaitForSelector(page playwright.Page, selector string) error {
	_, err := page.WaitForSelector(selector, playwright.PageWaitForSelectorOptions{
		Timeout: playwright.Float(30000),
	})
	return err
}

func WaitForSelectorAndClick(ctx context.Context, page playwright.Page, selector string) error {
	_, err := page.WaitForSelector(selector, playwright.PageWaitForSelectorOptions{
		Timeout: playwright.Float(30000),
	})
	if err != nil {
		slog.ErrorContext(ctx, "等待元素[Selector: %s]失败: %v", selector, err)
		return err
	}

	err = page.Click(selector)
	if err != nil {
		slog.ErrorContext(ctx, "点击按钮[Selector: %s]失败: %v", selector, err)
		return err
	}

	return nil
}

func NewPage(ctx context.Context, browser playwright.Browser, typeTag string) (playwright.Page, error) {
	var opt playwright.BrowserNewPageOptions
	if typeTag == LaoM || typeTag == ZhouLF || typeTag == JinZZ {
		opt = playwright.BrowserNewPageOptions{
			Viewport: &playwright.Size{
				Width:  375, // 手机屏幕宽度
				Height: 812, // 高度（如 iPhone 13 的竖屏尺寸）
			},
			IsMobile: playwright.Bool(true),
			// UserAgent: playwright.String("Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Mobile Safari/537.36"),
		}
	}

	return browser.NewPage(opt)
}
