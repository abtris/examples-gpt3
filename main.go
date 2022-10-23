package main

import (
	"context"
	"fmt"
	"os"

	gogpt "github.com/sashabaranov/go-gpt3"
)

func main() {
	gptToken := os.Getenv("GPT_TOKEN")
	c := gogpt.NewClient(gptToken)
	ctx := context.Background()

	prompt := "Say poem about ice cream"

	req := gogpt.CompletionRequest{
		Model:            "text-davinci-002",
		MaxTokens:        256,
		Temperature:      0.7,
		Prompt:           prompt,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
		BestOf:           1,
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		return
	}
	fmt.Printf("Prompt: %s\n\n", prompt)
	for _, choice := range resp.Choices {
		fmt.Println(choice.Text)
	}

}
