package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/fabiustech/anthropic"
)

var key *string

func init() {
	key = flag.String("key", "", "api key")
	flag.Parse()
}

func main() {
	var client = anthropic.NewClient(*key)
	var resp, err = client.NewCompletion(context.Background(), &anthropic.Request{
		Prompt:            anthropic.NewPromptFromString("Tell me a haiku about trees"),
		Model:             anthropic.Claude,
		MaxTokensToSample: 300,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Completion)
}
