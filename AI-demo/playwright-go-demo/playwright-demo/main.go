package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/playwright-community/playwright-go"
)

func main() {
	// 初始化Playwright
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("无法启动playwright: %v", err)
	}
	defer pw.Stop()

	// 启动浏览器
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false), // 设置为false可以看到浏览器界面
	})
	if err != nil {
		log.Fatalf("无法启动浏览器: %v", err)
	}
	defer browser.Close()

	// 创建一个新的页面
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("无法创建新页面: %v", err)
	}

	// 访问登录页面 - 将下面的URL替换为实际的登录页面
	fmt.Println("正在访问登录页面...")
	if _, err = page.Goto("http://47.122.117.74/signin", playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateNetworkidle,
	}); err != nil {
		log.Fatalf("无法访问登录页面: %v", err)
	}

	// 定位输入框并输入账号
	fmt.Println("正在输入账号...")
	if err = page.Locator("#email").Fill("461694355@qq.com"); err != nil {
		log.Fatalf("无法输入用户名: %v", err)
	}

	// 定位输入框并输入密码
	fmt.Println("正在输入密码...")
	if err = page.Locator("#password").Fill("Aptx4869000"); err != nil {
		log.Fatalf("无法输入密码: %v", err)
	}

	// 点击登录按钮
	fmt.Println("正在点击登录按钮...")
	// 定义选择器（推荐以下两种方式之一）
	//selector := "button.btn-primary:has-text('登录')" // 结合类名和文本定位
	// 或
	selector := "button[type='button'][tabindex='2']" // 结合 type 和 tabindex 属性定位

	// 点击操作（带强制点击选项）
	err = page.Click(selector)
	if err != nil {
		fmt.Printf("点击失败: %v\n", err)
		return
	}

	// 等待登录完成
	fmt.Println("正在等待登录完成...")
	if err = page.WaitForURL("http://47.122.117.74/apps", playwright.PageWaitForURLOptions{
		Timeout: playwright.Float(30000),
	}); err != nil {
		log.Fatalf("登录失败或超时: %v", err)
	}

	// 方法1：获取全部 Local Storage
	lsAll, _ := page.Evaluate(`() => JSON.stringify(window.localStorage)`)
	fmt.Printf("完整 Local Storage:\n%s\n", lsAll)

	// 获取所有cookies
	//fmt.Println("登录成功，正在获取Cookies...")
	//cookies, err := page.Context().Cookies()
	//if err != nil {
	//	log.Fatalf("无法获取cookies: %v", err)
	//}

	// 将cookies保存到文件
	cookiesJson, err := json.MarshalIndent(lsAll, "", "  ")
	if err != nil {
		log.Fatalf("无法序列化cookies: %v", err)
	}

	if err = os.WriteFile("cookies.json", cookiesJson, 0644); err != nil {
		log.Fatalf("无法保存cookies到文件: %v", err)
	}

	fmt.Println("Cookies已保存到cookies.json文件")

	// 展示一下已经成功登录的页面（可选）
	fmt.Println("登录成功，当前URL:", page.URL())

	// 截图保存（可选）
	if _, err = page.Screenshot(playwright.PageScreenshotOptions{
		Path:     playwright.String("login-success.png"),
		FullPage: playwright.Bool(true),
	}); err != nil {
		log.Printf("无法保存截图: %v", err)
	} else {
		fmt.Println("截图已保存为login-success.png")
	}
}
