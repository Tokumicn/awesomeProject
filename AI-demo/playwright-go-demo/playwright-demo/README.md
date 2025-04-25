# Playwright-Go 自动登录演示

这个项目演示如何使用 Playwright-Go 自动执行网站登录流程并保存 cookies。

## 功能

- 使用 Playwright-Go 自动化浏览器操作
- 访问登录页面
- 自动填写用户名和密码
- 点击登录按钮
- 保存登录后的 cookies 到 JSON 文件
- 保存登录成功的页面截图
- 读取保存的 cookies 并在新会话中使用

## 前置要求

- Go 1.17 或更高版本
- Playwright-Go 库

## 安装

1. 克隆此仓库
   ```
   git clone https://github.com/yourusername/playwright-demo.git
   cd playwright-demo
   ```

2. 安装依赖
   ```
   go mod tidy
   ```

3. 安装 Playwright 浏览器
   ```
   go run github.com/playwright-community/playwright-go/cmd/playwright install
   ```

## 使用方法

### 登录并保存 Cookies

1. 打开 `main.go` 文件，修改以下内容：
   - 登录页面的 URL (`https://example.com/login`)
   - 用户名输入框的选择器 (`#username`)
   - 密码输入框的选择器 (`#password`)
   - 登录按钮的选择器 (`button[type='submit']`)
   - 登录成功后的重定向 URL 模式 (`**/dashboard`)
   - 用户名和密码

2. 运行程序
   ```
   go run main.go
   ```

3. 程序将启动浏览器，执行登录流程，并将 cookies 保存到 `cookies.json` 文件。

### 使用保存的 Cookies

1. 先运行 `main.go` 以获取并保存 cookies。

2. 然后修改 `examples/cookies/load_cookies.go` 中的目标 URL，运行以下命令：
   ```
   go run examples/cookies/load_cookies.go
   ```

3. 程序将读取保存的 cookies，并使用它们访问需要登录的页面，无需再次登录。

## 注意事项

- 默认情况下，脚本会以非无头模式运行，这样你可以看到浏览器操作过程。如果需要无头模式运行，将 `Headless: playwright.Bool(false)` 改为 `Headless: playwright.Bool(true)`。
- 根据目标网站的实际情况，你可能需要调整选择器和等待时间。
- 确保以合法、道德的方式使用此脚本，不要用于未授权的访问或网站爬取。
- cookies 的有效期取决于目标网站的设置，过期后需要重新登录获取新的 cookies。 