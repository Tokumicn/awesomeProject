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

	// 读取保存的cookies
	fmt.Println("正在读取cookies...")
	cookiesData, err := os.ReadFile("cookies.json")
	if err != nil {
		log.Fatalf("无法读取cookies文件: %v", err)
	}

	// 解析cookies
	var storedCookies []playwright.Cookie
	if err = json.Unmarshal(cookiesData, &storedCookies); err != nil {
		log.Fatalf("无法解析cookies数据: %v", err)
	}

	// 转换为OptionalCookie格式
	cookies := make([]playwright.OptionalCookie, 0, len(storedCookies))
	for _, c := range storedCookies {
		domain := c.Domain
		path := c.Path
		expires := c.Expires
		httpOnly := c.HttpOnly
		secure := c.Secure

		cookie := playwright.OptionalCookie{
			Name:     c.Name,
			Value:    c.Value,
			Domain:   &domain,
			Path:     &path,
			Expires:  &expires,
			HttpOnly: &httpOnly,
			Secure:   &secure,
		}

		// 仅当SameSite不为空时添加
		if c.SameSite != nil {
			sameSite := *c.SameSite
			cookie.SameSite = &sameSite
		}

		cookies = append(cookies, cookie)
	}

	// 启动浏览器
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false), // 设置为false可以看到浏览器界面
	})
	if err != nil {
		log.Fatalf("无法启动浏览器: %v", err)
	}
	defer browser.Close()

	// 创建一个新的浏览器上下文
	context, err := browser.NewContext()
	if err != nil {
		log.Fatalf("无法创建浏览器上下文: %v", err)
	}

	// 添加cookies到上下文
	if err = context.AddCookies(cookies); err != nil {
		log.Fatalf("无法添加cookies: %v", err)
	}
	fmt.Println("Cookies已加载")

	// 创建一个新的页面
	page, err := context.NewPage()
	if err != nil {
		log.Fatalf("无法创建新页面: %v", err)
	}

	// 访问需要登录的页面 - 应该直接展示已登录状态
	// 将下面的URL替换为实际的URL
	fmt.Println("正在访问需要登录的页面...")
	if _, err = page.Goto("https://example.com/dashboard", playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateNetworkidle,
	}); err != nil {
		log.Fatalf("无法访问页面: %v", err)
	}

	// 检查是否成功登录
	fmt.Println("检查登录状态...")
	fmt.Println("当前URL:", page.URL())

	// 截图保存（可选）
	if _, err = page.Screenshot(playwright.PageScreenshotOptions{
		Path:     playwright.String("logged-in-with-cookies.png"),
		FullPage: playwright.Bool(true),
	}); err != nil {
		log.Printf("无法保存截图: %v", err)
	} else {
		fmt.Println("截图已保存为logged-in-with-cookies.png")
	}

	// 等待用户手动关闭
	fmt.Println("按回车键关闭浏览器...")
	fmt.Scanln()
}
