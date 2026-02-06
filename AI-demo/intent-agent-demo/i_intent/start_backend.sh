#!/bin/bash

# AI助手 - 后端服务启动脚本 

echo "🖥️ 启动AI助手后端服务..."

cd "$(dirname "$0")/i_intent"

echo "🚀 启动后端服务 (端口5050)..."
go run main.go
