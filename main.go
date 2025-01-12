package main

import (
	"bufio"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/ollama/ollama/api"
)

// showLoadingWheel displays a loading spinner in the terminal.
func showLoadingWheel(done chan bool) {
	spinner := []string{
		"▪▫▫▫",
		"▫▪▫▫",
		"▫▫▪▫",
		"▫▫▫▪",
		"▫▫▫▫",
	}
	i := 0
	for {
		select {
		case <-done:
			fmt.Print("\r")
			return
		default:
			i = (i + 1) % len(spinner)
			fmt.Printf("\rGenerating response... %s", spinner[i])
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func main() {
	fmt.Println("Welcome to the Ollama Chat! Type 'exit' to quit.")

	// Scanner for user input
	scanner := bufio.NewScanner(os.Stdin)

	// Create the Ollama API client
	baseURL, _ := url.Parse("http://localhost:11434") // Replace with your Ollama API base URL
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

		// Channel to signal when the loading wheel should stop
		done := make(chan bool)
		go showLoadingWheel(done)

		// Send the request to the Ollama API
		var response string
		err := client.Generate(context.Background(), &api.GenerateRequest{
			Model:  "llama3.2",
			Prompt: userInput,
		}, func(cr api.GenerateResponse) error {
			if cr.Response != "" {
				response += cr.Response // Append the response for multi-part messages
			}
			return nil
		})

		// Stop the loading spinner
		done <- true

		if err != nil {
			fmt.Printf("Error sending request: %v\n", err)
			continue
		}

		// Print the response
		if response == "" {
			fmt.Println("Ollama: No response received.")
		} else {
			fmt.Printf("Ollama: %s\n", response)
		}
	}
}
