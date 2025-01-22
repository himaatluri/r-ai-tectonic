package main

import (
	"bufio"
	"context"
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
		fmt.Print("⮑ : ")
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
			fmt.Println("Document based inference comming soon.. to analyze this file: ", filename)
		}

		done := make(chan bool)
		go helpers.ShowLoadingWheel(done)

		// Send the request to the Ollama API
		var response string
		err := client.Generate(context.Background(), &api.GenerateRequest{
			Model:  "phi4",
			Prompt: userInput,
		}, func(cr api.GenerateResponse) error {
			if cr.Response != "" {
				response += cr.Response
			}
			return nil
		})

		done <- true

		if err != nil {
			fmt.Printf("Error sending request: %v\n", err)
			continue
		}

		if response == "" {
			fmt.Println("Ollama: No response received.")
		} else {
			fmt.Print("Ollama: \n")
			helpers.StreamResponse(response)
		}
	}
}
