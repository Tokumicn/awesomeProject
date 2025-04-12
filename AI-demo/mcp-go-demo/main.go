package main

import (
	"fmt"
	"log"

	"github.com/ThinkInAIXYZ/go-mcp/protocol"
	"github.com/ThinkInAIXYZ/go-mcp/server"
	"github.com/ThinkInAIXYZ/go-mcp/transport"
)

func main() {
	// 创建传输服务器（本例中使用 SSE）
	transportServer, err := transport.NewSSEServerTransport("")
	if err != nil {
		log.Fatalf("创建传输服务器失败: %v", err)
	}

	// 使用传输创建 MCP 服务器
	mcpServer, err := server.NewServer(transportServer,
		// 设置服务器实现信息
		server.WithServerInfo(protocol.Implementation{
			Name:    "示例 MCP 服务器",
			Version: "1.0.0",
		}),
	)
	if err != nil {
		log.Fatalf("创建 MCP 服务器失败: %v", err)
	}

	// 注册工具处理器
	mcpServer.RegisterTool(&protocol.Tool{
		Name:        "example_tool",
		Description: "example_tool",
		InputSchema: make(map[string]interface{}),
	}, func(req *protocol.CallToolRequest) (*protocol.CallToolResult, error) {
		fmt.Printf("工具被调用，参数: %+v\n", req.Arguments)
		return protocol.NewCallToolResult([]protocol.Content{protocol.TextContent{
			Annotated: protocol.Annotated{},
			Type:      "text",
			Text:      "测试",
		}}, true), nil
	})

	if err = mcpServer.Start(); err != nil {
		log.Fatalf("启动 MCP 服务器失败: %v", err)
		return
	}
}
