package main

import (
	"context"
	"playground/modal"
	"playground/protocol"

	"github.com/aws/aws-lambda-go/lambda"
	"google.golang.org/protobuf/proto"
)

func HandleRequest(ctx context.Context, req modal.DataCollectorRequest) (modal.DataCollectorResponse, error) {
	res := modal.DataCollectorResponse{}
	job := &protocol.PanopticJob{}
	proto.Unmarshal(req.SerializedJob, job)

	marshall, err := proto.Marshal(job)
	if err != nil {
		return res, err
	}

	res.SerializedJob = marshall
	return res, nil
}

func main() {
	lambda.Start(HandleRequest)
}
