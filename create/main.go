package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
)

func main() {
	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-west-1"),
	)

	if err != nil {
		panic(err)
	}

	client := lambda.NewFromConfig(cfg)

	funcName := "testing_func"
	aws_role := "arn:aws:iam::213288384225:role/service-role/test_ddog_logging-role-8qnsddqu"
	aws_container_image := "213288384225.dkr.ecr.us-west-1.amazonaws.com/hello-world:latest"
	timeout := int32(300)

	_, err = client.CreateFunction(ctx, &lambda.CreateFunctionInput{
		FunctionName: &funcName,
		Role:         &aws_role,
		Code: &types.FunctionCode{
			ImageUri: &aws_container_image,
		},
		Timeout:     &timeout,
		Description: &funcName,
		PackageType: types.PackageTypeImage,
	})

	if err != nil {
		panic(err)
	}
}
