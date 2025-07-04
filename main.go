package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
}

func HandleRequest(ctx context.Context, event *MyEvent) (string, error) {
	if event == nil {
		return "", fmt.Errorf("received nil event")
	}
	return fmt.Sprintf("Hello, World!"), nil
}

func main() {
	lambda.Start(HandleRequest)
}

