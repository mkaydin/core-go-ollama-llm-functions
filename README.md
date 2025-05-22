---

# core-go-ollama-llm-functions

A collection of Go-based examples demonstrating how to integrate Large Language Models (LLMs) with function calling capabilities using [Ollama](https://github.com/ollama/ollama).

## Overview

This repository showcases various implementations of function calling with LLMs in Go, leveraging Ollama's local model serving capabilities. It includes examples such as:

* Structured JSON outputs
* Tool usage and integration
* Reasoning and inference tasks
* Retrieval-Augmented Generation (RAG) workflows
* SQLite-backed data interactions
* HTTP client-server communication

Each example is designed to illustrate specific aspects of LLM function calling, providing a practical reference for developers.

## Prerequisites

* Go 1.20 or higher
* [Ollama](https://github.com/ollama/ollama) installed and running locally
* Required LLM models downloaded via Ollama (e.g., Llama 3.1)

## Getting Started

1. **Clone the repository:**

   ```bash
   git clone https://github.com/mkaydin/core-go-ollama-llm-functions.git
   cd core-go-ollama-llm-functions
   ```



2. **Install dependencies:**

   ```bash
   go mod tidy
   ```



3. **Run an example:**

   Navigate to the desired example directory and execute the Go file. For instance:

   ```bash
   cd structured_output_v1
   go run main.go
   ```



## Project Structure

The repository is organized into the following directories:

* `structured_output_v1/` - Basic structured output example
* `structured_output_v2/` - Advanced structured output with nested data
* `tools_output/` - Demonstrates tool usage within LLM responses
* `reasoning/` - Examples focusing on reasoning capabilities
* `rag_output/` - Retrieval-Augmented Generation implementation
* `mcp_sqlite_docker/` - SQLite integration with Docker support
* `mcp-curl-client/` - HTTP client example using curl
* `mcp-curl-server/` - HTTP server handling LLM interactions
* `json_output/` - JSON output formatting examples

## Usage

Each example is self-contained. To run a specific example:

1. Ensure Ollama is running and the necessary model is available:

   ```bash
   ollama run llama3.1
   ```



2. Navigate to the example directory:

   ```bash
   cd <example_directory>
   ```



3. Execute the Go program:

   ```bash
   go run main.go
   ```



Replace `<example_directory>` with the desired example folder name.

## Contributing

Contributions are welcome! If you have suggestions or improvements, feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

For more information on Ollama and its capabilities, visit the [official Ollama GitHub repository](https://github.com/ollama/ollama).

---
