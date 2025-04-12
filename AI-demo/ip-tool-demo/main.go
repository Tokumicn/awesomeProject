package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"io"
	"log/slog"
	"net"
	"net/http"
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
		"ip-mcp",
		"1.0.0",
	)

	// Add tool
	tool := mcp.NewTool("ip_query",
		mcp.WithDescription("query geo location of an IP address"),
		mcp.WithString("ip",
			mcp.Required(),
			mcp.Description("IP address to query"),
		),
	)

	// Add tool handler
	s.AddTool(tool, ipQueryHandler)

	// Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

func ipQueryHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ip, ok := request.Params.Arguments["ip"].(string)
	if !ok {
		return nil, errors.New("ip must be a string")
	}

	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		slog.ErrorContext(ctx, "invalid IP address: %s", ip)
		return nil, errors.New("invalid IP address")
	}

	resp, err := http.Get("https://ip.rpcx.io/api/ip?ip=" + ip)
	if err != nil {
		slog.ErrorContext(ctx, "Error fetching IP information: %v", err)
		return nil, fmt.Errorf("Error fetching IP information: %v", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.ErrorContext(ctx, "Error reading response body: %v", err)
		return nil, fmt.Errorf("Error reading response body: %v", err)
	}

	return mcp.NewToolResultText(string(data)), nil
}
