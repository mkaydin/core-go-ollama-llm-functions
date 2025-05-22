package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/ollama/ollama/api"
)

func main() {
	ctx := context.Background()

	ollamaUrl := os.Getenv("OLLAMA_HOST")
	if ollamaUrl == "" {
		ollamaUrl = "http://localhost:11434"
	}

	model := os.Getenv("LLM")
	if model == "" {
		model = "deepseek-r1:7b"
	}

	fmt.Println("🌍", ollamaUrl, "📕", model)

	url, _ := url.Parse(ollamaUrl)
	client := api.NewClient(url, http.DefaultClient)

	systemInstructions, err := os.ReadFile("system-instructions.md")
	if err != nil {
		log.Fatal("😡:", err)
	}

	irisInstructions, err := os.ReadFile("iris-instructions.md")
	if err != nil {
		log.Fatal("😡:", err)
	}

	irisDatabase, err := os.ReadFile("iris-database.xml")
	if err != nil {
		log.Fatal("😡:", err)
	}

	userContent := `Using the above information and the below information,
			Given a specimen with:
			- Petal width: 1,9 cm
			- Petal length: 5,1 cm
			- Sepal width: 2,7 cm
			- Sepal length: 5,8 cm
			What is the species of the iris?`

	messages := []api.Message{
		{Role: "system", Content: string(systemInstructions)},
		{Role: "system", Content: string(irisInstructions)},
		{Role: "system", Content: "# Iris Database\n" + string(irisDatabase)},
		{Role: "user", Content: userContent},
	}

	stream := true

	req := &api.ChatRequest{
		Model:    model,
		Messages: messages,
		Options: map[string]interface{}{
			"temperature":    0.0,
			"repeat_last_n":  2,
			"repeat_penalty": 2.2,
			"top_k":          10,
			"top_p":          0.5,
		},
		KeepAlive: &api.Duration{Duration: 1 * time.Minute},
		Stream:    &stream,
	}

	var fullResponse string

	respFunc := func(resp api.ChatResponse) error {
		fmt.Print(resp.Message.Content)
		fullResponse += resp.Message.Content
		return nil
	}

	errChat := client.Chat(ctx, req, respFunc)
	if errChat != nil {
		log.Fatal("😡:", errChat)
	}

	// Save to markdown file
	outputFile := "llm-response.md"
	err = os.WriteFile(outputFile, []byte(fullResponse), 0644)
	if err != nil {
		log.Fatal("💥 Failed to save response to file:", err)
	}

	fmt.Println("\n✅ LLM response saved to", outputFile)
}
