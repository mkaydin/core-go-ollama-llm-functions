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
	FALSE = false
	TRUE  = true
)

func main() {
	ctx := context.Background()

	var ollamaRawUrl string
	if ollamaRawUrl = os.Getenv("OLLAMA_HOST"); ollamaRawUrl == "" {
		ollamaRawUrl = "http://localhost:11434"
	}

	url, _ := url.Parse(ollamaRawUrl)

	client := api.NewClient(url, http.DefaultClient)

	// define schema for a structured output
	// ref: https://ollama.com/blog/structured-outputs
	schema := map[string]any{
		"type": "object",
		"properties": map[string]any{
			"scientific_name": map[string]any{
				"type": "string",
			},
			"main_species": map[string]any{
				"type": "string",
			},
			"average_length": map[string]any{
				"type": "number",
			},
			"average_lifespan": map[string]any{
				"type": "number",
			},
			"average_weight": map[string]any{
				"type": "number",
			},
			"countries": map[string]any{
				"type": "array",
				"items": map[string]any{
					"type": "string",
				},
			},
		},
		"required": []string{"scientific_name", "main_species", "average_length", "average_lifespan", "average_weight", "countries"},
	}

	// convert schema to json with pretty print
	jsonModel, err := json.MarshalIndent(schema, "", "  ")

	fmt.Println(string(jsonModel))
	if err != nil {
		log.Fatalln("😡", err)
	}

	data := `Information about the chicken:
			- scientific_name: Gallus gallus
			- main_species: Poultry
			- average_length: 1.5 to 1.75 meters
			- average_weight: 5 to 7 kilograms
			- average_lifespan: 10 to 20 years
			- countries: ["China", "Iran", "India", "Egypt", "Turkey"]
			`

	userContent := "tell me about chicken"

	messages := []api.Message{
		{Role: "system", Content: data},
		{Role: "user", Content: userContent},
	}

	req := &api.ChatRequest{
		Model:    "gemma3:4b",
		Messages: messages,
		Options: map[string]interface{}{
			"temperature":   0.0,
			"repeat_last_n": 2,
		},
		Stream: &FALSE,
		Format: json.RawMessage(jsonModel),
	}

	answer := ""
	err = client.Chat(ctx, req, func(resp api.ChatResponse) error {
		answer = resp.Message.Content
		return nil
	})

	if err != nil {
		log.Fatalln("😡", err)
	}
	fmt.Println(answer)
	fmt.Println()
}
