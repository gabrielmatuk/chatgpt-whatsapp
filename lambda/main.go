package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gabrielmatuk/wppgpt/src/chatgpt"
	"github.com/gabrielmatuk/wppgpt/src/config"
)

func init() {
	config.LoadingEnv()
	fmt.Println("Loading ENV variables.")
}

func main() {
	lambda.Start(chatgpt.Process)
	fmt.Println("Message send!")
}
