package main

import (
	"context"
	"encoding/json"
	"fmt"
	"playground/modal"
	"playground/protocol"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"google.golang.org/protobuf/proto"
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

	job := &protocol.PanopticJob{}
	job.JobId = "test_id"
	marshall, err := proto.Marshal(job)
	if err != nil {
		panic(err)
	}

	req := modal.DataCollectorRequest{
		SerializedJob: marshall,
	}
	payload, err := json.Marshal(req)
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

	var collectorResponse modal.DataCollectorResponse

	err = json.Unmarshal(res.Payload, &collectorResponse)
	if err != nil {
		panic(err)
	}

	resJob := &protocol.PanopticJob{}
	err = proto.Unmarshal(collectorResponse.SerializedJob, resJob)

	fmt.Println(resJob.String())
	if err != nil {
		panic(err)
	}
}
