package get_my_ip

import (
	"context"
	"errors"
	"github.com/mark3labs/mcp-go/mcp"
	"log/slog"
	"net"
)

func GetMyIPHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	ip := getIPAddress()

	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		slog.ErrorContext(ctx, "invalid IP address")
		return nil, errors.New("invalid IP address")
	}

	return mcp.NewToolResultText(ip), nil
}

func getIPAddress() string {
	interfaces, _ := net.Interfaces()
	for _, iface := range interfaces {
		// 过滤非启用或回环接口
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, _ := iface.Addrs()
		for _, addr := range addrs {
			if ipNet, ok := addr.(*net.IPNet); ok && ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}
	return ""
}
