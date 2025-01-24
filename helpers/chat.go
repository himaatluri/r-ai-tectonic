package helpers

import (
	"context"
	"fmt"

	"github.com/ollama/ollama/api"
)

func InvokeChat(client *api.Client, input string) {
	done := make(chan bool)
	go ShowLoadingWheel(done)

	// Send the request to the Ollama API
	var response string
	err := client.Generate(context.Background(), &api.GenerateRequest{
		Model:  "phi4",
		Prompt: input,
	}, func(cr api.GenerateResponse) error {
		if cr.Response != "" {
			response += cr.Response
		}
		return nil
	})

	done <- true

	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
	}

	if response == "" {
		fmt.Println("Ollama: No response received.")
	} else {
		fmt.Print("Ollama: \n")
		StreamResponse(response)
	}
}
