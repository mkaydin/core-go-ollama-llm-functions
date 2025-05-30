package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/mcp"
)

func main() {

	// create a new MCP client with the server data connection
	mcpClient, err := client.NewStdioMCPClient(
		"docker",
		[]string{}, // environment variables
		"run",
		"--rm",
		"-i",
		"mcp-curl",
	)
	if err != nil {
		log.Fatalf(":C Failed to create client: %v", err)
	}
	defer mcpClient.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// define and initialize the MCP request
	fmt.Println("Initializing mcp client...")
	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "mcp-curl client <3",
		Version: "1.0.0",
	}

	// initialize the MCP client and connect to the server
	initResult, err := mcpClient.Initialize(ctx, initRequest)
	if err != nil {
		log.Fatalf(":CFailed to initialize: %v", err)
	}
	fmt.Printf(
		"Initialized with server: %s %s\n\n",
		initResult.ServerInfo.Name,
		initResult.ServerInfo.Version,
	)

	// get the list of tools from the server
	fmt.Println("Available tools...")
	toolsRequest := mcp.ListToolsRequest{}
	tools, err := mcpClient.ListTools(ctx, toolsRequest)
	if err != nil {
		log.Fatalf(":C Failed to list tools: %v", err)
	}

	// display the list of tools
	for _, tool := range tools.Tools {
		fmt.Printf("- %s: %s\n", tool.Name, tool.Description)
		fmt.Println("Arguments: ", tool.InputSchema.Properties)
	}
	fmt.Println()

	// prepare the call of the tool "use_curl"
	// to fetch the connect of a web page
	fmt.Println("Calling use_curl")
	fetchRequest := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	fetchRequest.Params.Name = "use_curl"
	fetchRequest.Params.Arguments = map[string]interface{}{
		"url": "https://raw.githubusercontent.com/docker-sa/01-build-image/refs/heads/main/main.go",
	}

	// call the tool
	result, err := mcpClient.CallTool(ctx, fetchRequest)
	if err != nil {
		log.Fatalf("Failed to call the tool: %v", err)
	}

	// display the text content of result
	fmt.Println("Content of the page: ")
	fmt.Println(result.Content[0].(map[string]interface{})["text"])
}
