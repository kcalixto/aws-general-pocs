package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-xray-sdk-go/lambda"
)

func Handler(ctx context.Context, event events.SQSEvent) (res events.SQSEventResponse, err error) {
	var failures []events.SQSBatchItemFailure
	for _, record := range event.Records {
		
	}
}

func main() {
	lambda.Start(Handler)
}
