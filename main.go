package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	gogpt "github.com/sashabaranov/go-gpt3"
)

func main() {
	gptToken := os.Getenv("GPT_TOKEN")
	c := gogpt.NewClient(gptToken)
	ctx := context.Background()
	prompt := flag.String("prompt", "Say poem about ice cream", "prompt input for GPT-3")
	flag.Parse()
	req := gogpt.CompletionRequest{
		Model:            "text-davinci-002",
		MaxTokens:        256,
		Temperature:      0.7,
		Prompt:           *prompt,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
		BestOf:           1,
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		return
	}
	fmt.Printf("Prompt: %s\n\n", *prompt)
	for _, choice := range resp.Choices {
		fmt.Println(choice.Text)
	}

}
