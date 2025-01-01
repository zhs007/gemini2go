package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/genai"
)

func printResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
			}
		}
	}
	fmt.Println("---")
}

func main() {
	ctx := context.Background()
	apiKey := os.Getenv("GEMINI_API_KEY")
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGoogleAI,
	})
	if err != nil {
		log.Fatal(err)
	}

	parts := []*genai.Part{
		{Text: "Write a story about a magic backpack in Chinese."},
	}
	result, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash-exp", []*genai.Content{{Parts: parts}}, nil)
	printResponse(result)
}
