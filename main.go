package main

import (
	"context"
	"fmt"
	"playground/modal"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, name modal.MyEvent) (string, error) {
	return fmt.Sprintf("Hello %s!", name.Name), nil
}

func main() {
	lambda.Start(HandleRequest)
}
