package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-west-1"),
	)

	if err != nil {
		panic(err)
	}

	client := lambda.NewFromConfig(cfg)

	ctx := context.Background()
	funcName := "testing_func"

	_, err = client.DeleteFunction(ctx, &lambda.DeleteFunctionInput{
		FunctionName: &funcName,
	})

	if err != nil {
		panic(err)
	}
}
