package main

import (
	"context"
	"fmt"
	"github.com/ollama/ollama/api"
	"github.com/presbrey/ollamafarm"
	"log"
	"time"
)

var model = "qwen2:0.5b"

func main() {
	farm := ollamafarm.New()
	// Register Ollama servers in the same group with different priorities
	farm.RegisterURL("http://192.168.10.175:11434", &ollamafarm.Properties{Group: "3060", Priority: 1})
	farm.RegisterURL("http://192.168.10.192:11434", &ollamafarm.Properties{Group: "4090", Priority: 2})
	var ollama *ollamafarm.Ollama
	retryCount, maxRetries := 0, 5
	for ollama == nil && retryCount < maxRetries {
		ollama = farm.First(&ollamafarm.Where{Model: model})
		time.Sleep(time.Second)
		retryCount++
	}
	fmt.Println(retryCount)
	if ollama != nil {
		// Perform a Chat call
		req := &api.ChatRequest{
			Model: model,
			Messages: []api.Message{
				{Role: "user", Content: "你好"},
			},
		}

		err := ollama.Client().Chat(context.Background(), req, func(resp api.ChatResponse) error {
			fmt.Print(resp.Message.Content)
			return nil
		})

		if err != nil {
			log.Fatalf("Chat error: %v", err)
		}
		fmt.Println()
	}
	// Get model counts
	modelCounts := farm.ModelCounts(nil)
	fmt.Printf("Available models: %v\n", modelCounts)
}
