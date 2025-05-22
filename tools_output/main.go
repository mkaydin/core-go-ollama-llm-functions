package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/ollama/ollama/api"
)

var (
	TRUE  = true
	FALSE = false
)

func main() {
	ctx := context.Background()

	var ollamaRawUrl string
	if ollamaRawUrl = os.Getenv("OLLAMA_HOST"); ollamaRawUrl == "" {
		ollamaRawUrl = "http://localhost:11434"
	}

	var toolsLLM string
	if toolsLLM = os.Getenv("TOOLS_LLM"); toolsLLM == "" {
		toolsLLM = "qwen2.5:1.5b"
	}

	url, _ := url.Parse(ollamaRawUrl)

	client := api.NewClient(url, http.DefaultClient)

	// define some tools
	helloTool := map[string]any{
		"type": "function",
		"function": map[string]any{
			"name":        "hello",
			"description": "Sya hello to a given person with his name",
			"parameters": map[string]any{
				"type": "object",
				"properties": map[string]any{
					"name": map[string]any{
						"type":        "string",
						"description": "the name of the person",
					},
				},
				"required": []string{"name"},
			},
		},
	}

	addNumbersTool := map[string]any{
		"type": "function",
		"function": map[string]any{
			"name":        "add_numbers",
			"description": "Add two numbers",
			"parameters": map[string]any{
				"type": "object",
				"properties": map[string]any{
					"number1": map[string]any{
						"type":        "number",
						"description": "The first number",
					},
					"number2": map[string]any{
						"type":        "number",
						"description": "The second number",
					},
				},
				"required": []string{"number1", "number2"},
			},
		},
	}

	tools := []any{helloTool, addNumbersTool}

	// transform tools to json
	jsonTools, _ := json.Marshal(tools)

	var toolsList api.Tools
	jsonErr := json.Unmarshal(jsonTools, &toolsList)
	if jsonErr != nil {
		log.Fatalln(":C", jsonErr)
	}

	// Prompt construction
	messages := []api.Message{
		{Role: "user", Content: "Say hello to Bob"},
		{Role: "User", Content: "add 31 to 69"},
		{Role: "user", Content: "Say hello to Sarah"},
	}

	req := &api.ChatRequest{
		Model:    toolsLLM,
		Messages: messages,
		Options: map[string]interface{}{
			"temperature":   0.0,
			"repeat_last_n": 2,
		},
		Tools:  toolsList,
		Stream: &FALSE,
	}

	err := client.Chat(ctx, req, func(resp api.ChatResponse) error {
		//fmt.Println("🖐️", resp.Message.ToolCalls)

		for _, toolCall := range resp.Message.ToolCalls {
			fmt.Println(toolCall.Function.Name, toolCall.Function.Arguments)
		}

		return nil
	})

	if err != nil {
		log.Fatalln(":C", err)
	}

	fmt.Println()
	fmt.Println("O_o", ollamaRawUrl, toolsLLM)
}
