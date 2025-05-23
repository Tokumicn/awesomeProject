package ip2location

import (
	"context"
	"errors"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"io"
	"log/slog"
	"net"
	"net/http"
)

func IPQueryHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
