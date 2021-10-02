package main

import (
	"context"
	"encoding/json"
	"fmt"
	"playground/modal"

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

	me := modal.MyEvent{
		Name: "world",
	}
	payload, err := json.Marshal(me)
	if err != nil {
		panic(err)
	}

	res, err := client.Invoke(ctx, &lambda.InvokeInput{
		FunctionName: &funcName,
		Payload:      payload,
	})

	if err != nil {
		panic(err)
	}

	wording := string(res.Payload)
	fmt.Println(wording)
}
