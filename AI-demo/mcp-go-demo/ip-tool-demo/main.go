package main

import (
	"ai-demo/mcp-go-demo/ip-tool-demo/ip2location"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"log/slog"
	"os"
)

func main() {
	l := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))
	// slog.SetLogLoggerLevel(slog.LevelDebug)
	slog.SetDefault(l)

	// Create MCP server
	s := server.NewMCPServer(
		"ip-tools-mcp",
		"1.0.0",
	)

	// Add tool
	toolIP2Location := mcp.NewTool("ip_query_location",
		mcp.WithDescription("query geo location of an IP address"),
		mcp.WithDescription("查询IP地理位置"),
		mcp.WithString("ip",
			mcp.Required(),
			mcp.Description("IP address to query"),
		),
	)

	//// Add tool
	//toolGetMyIP := mcp.NewTool("get_my_ip",
	//	mcp.WithDescription("get my IP address"),
	//	mcp.WithDescription("获取我的IP地址"),
	//	mcp.WithDescription("获取本机IP"),
	//)

	// Add tool handler
	s.AddTool(toolIP2Location, ip2location.IPQueryHandler)

	//// Add tool handler
	//s.AddTool(toolGetMyIP, get_my_ip.GetMyIPHandler)

	// Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
