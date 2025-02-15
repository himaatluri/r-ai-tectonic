package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/himaatluri/r-ai-tectonic/helpers"
	"github.com/ollama/ollama/api"
)

func main() {
	fmt.Println("Welcome to the Ollama Chat! Type 'exit' to quit.")

	scanner := bufio.NewScanner(os.Stdin)

	baseURL, _ := url.Parse("http://localhost:11434")
	client := api.NewClient(baseURL, &http.Client{})

	for {
		fmt.Print("🤖 : ")
		if !scanner.Scan() {
			fmt.Println("Error reading input. Exiting.")
			break
		}

		userInput := scanner.Text()
		if userInput == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		if filename, matchedType := helpers.DetectDocInference(userInput); matchedType {
			fmt.Println("🗎 Document based inference")
			fileInput, err := os.ReadFile(filename)
			if err != nil {
				fmt.Println(err)
			}
			helpers.InvokeChat(client, string(fileInput))
			continue
		}

		helpers.InvokeChat(client, userInput)
	}
}
