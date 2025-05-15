package main

import (
	"context"
	"fmt"
	"github.com/playwright-community/playwright-go"
	"gold/playwright-gold/caibai"
	"gold/playwright-gold/common"
	"gold/playwright-gold/jinzhizun"
	"gold/playwright-gold/laofengxiang"
	"gold/playwright-gold/laomiao"
	"gold/playwright-gold/zhoudafu"
	"gold/playwright-gold/zhouliufu"
	"gold/playwright-gold/zhoushengsheng"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

var imagePath string

func init() {
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	binDir := filepath.Dir(exePath)
	fmt.Println("[GOLD] bin work dir: ", binDir)
	imagePath = fmt.Sprintf("%s/images/%s", binDir, time.Now().Format(time.DateOnly))
	// 文件夹不存在则创建
	err = os.MkdirAll(imagePath, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// 日志配置
	l := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))
	// slog.SetLogLoggerLevel(slog.LevelDebug)
	slog.SetDefault(l)

	// playwright安装依赖
	err = playwright.Install()
	if err != nil {
		panic(err)
	}
}

func main() {
	ctx := context.TODO()

	// 初始化Playwright
	pw, err := playwright.Run()
	if err != nil {
		slog.ErrorContext(ctx, "无法启动playwright: %v", err)
		return
	}
	defer pw.Stop()

	// 启动浏览器
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(!common.ShowChrome), // 设置为false可以看到浏览器界面
	})
	if err != nil {
		slog.ErrorContext(ctx, "无法启动浏览器: %v", err)
		return
	}
	defer browser.Close()

	tags := []string{
		common.ZhouSS,
		common.LaoFX,
		common.ZhouLF,
		common.ZhouDF,
		common.LaoM,
		common.CaiBai,
		// common.JinZZ,
	}

	for _, tag := range tags {
		// 创建一个新的页面
		page, err := common.NewPage(ctx, browser, tag)
		if err != nil {
			slog.ErrorContext(ctx, "创建页面错误: %v", err)
			continue
		}

		switch tag {
		case common.ZhouSS:
			err = zhoushengsheng.GetZhouShSh(ctx, page)
			if err != nil {
				goto closeTag
			}
		case common.LaoFX:
			err = laofengxiang.GetLaoFengXiang(ctx, page)
			if err != nil {
				goto closeTag
			}
		case common.ZhouLF:
			err = zhouliufu.GetZhouLiuFu(ctx, page)
			if err != nil {
				goto closeTag
			}
		case common.ZhouDF:
			err = zhoudafu.GetZhouDaFu(ctx, page)
			if err != nil {
				goto closeTag
			}
		case common.LaoM:
			err = laomiao.GetLaoMiao(ctx, page)
			if err != nil {
				goto closeTag
			}
		case common.CaiBai:
			err = caibai.GetCaiBai(ctx, page)
			if err != nil {
				goto closeTag
			}
		case common.JinZZ:
			err = jinzhizun.GetJinZhiZun(ctx, page)
			if err != nil {
				goto closeTag
			}
		default: // 不处理
			goto closeTag
		}

		// 截图保存
		if _, err = page.Screenshot(playwright.PageScreenshotOptions{
			Path:     playwright.String(fmt.Sprintf("%s/%s-gold-info-%d.png", imagePath, tag, time.Now().Unix())),
			FullPage: playwright.Bool(true),
		}); err != nil {
			slog.ErrorContext(ctx, "无法保存截图: %v", err)
		} else {
			slog.InfoContext(ctx, "截图已保存为login-success.png")
		}

	closeTag:
		page.Close() // 关闭页面
	}
	slog.InfoContext(ctx, "今日 ", time.Now().Format(time.DateOnly), " 所有页面截图已保存")
}
