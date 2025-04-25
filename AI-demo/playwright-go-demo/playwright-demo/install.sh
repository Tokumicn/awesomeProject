#!/bin/bash

echo "正在设置 Playwright-Go 环境..."

# 确保依赖已安装
echo "正在更新Go模块依赖..."
go mod tidy

# 安装Playwright浏览器
echo "正在安装Playwright浏览器（这可能需要几分钟）..."
go run github.com/playwright-community/playwright-go/cmd/playwright install

echo "安装完成！"
echo "使用方法："
echo "1. 修改 main.go 中的URL和选择器以匹配您的目标网站"
echo "2. 运行 'go run main.go' 登录并保存cookies"
echo "3. 运行 'go run examples/cookies/load_cookies.go' 使用保存的cookies访问网站" 