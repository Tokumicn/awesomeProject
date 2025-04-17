package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/ThinkInAIXYZ/go-mcp/client"
	"github.com/ThinkInAIXYZ/go-mcp/protocol"
	"github.com/ThinkInAIXYZ/go-mcp/transport"
)

func main() {
	// Create transport client (using SSE in this example)
	transportClient, err := transport.NewSSEClientTransport("http://127.0.0.1:8080/sse")
	if err != nil {
		log.Fatalf("Failed to create transport client: %v", err)
	}

	// Create MCP client using transport
	mcpClient, err := client.NewClient(transportClient, client.WithClientInfo(protocol.Implementation{
		Name:    "示例 MCP 客户端",
		Version: "1.0.0",
	}))
	if err != nil {
		log.Fatalf("Failed to create MCP client: %v", err)
	}
	defer mcpClient.Close()

	// List available tools
	toolsResult, err := mcpClient.ListTools(context.Background())
	if err != nil {
		log.Fatalf("Failed to list tools: %v", err)
	}
	b, _ := json.Marshal(toolsResult.Tools)
	fmt.Printf("Available tools: %+v\n", string(b))

	// Call tool
	callResult, err := mcpClient.CallTool(
		context.Background(),
		protocol.NewCallToolRequest("current time", map[string]interface{}{
			"timezone": "UTC",
		}))
	if err != nil {
		log.Fatalf("Failed to call tool: %v", err)
	}
	b, _ = json.Marshal(callResult)
	fmt.Printf("Tool call result: %+v\n", string(b))
}
