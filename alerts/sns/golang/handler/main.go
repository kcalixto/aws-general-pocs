package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context, event events.SNSEvent) error {
	for _, record := range event.Records {
		snsRecord := record.SNS
		fmt.Printf("Message received from SNS: %s\n", snsRecord.Message)
	}

	return nil
}

func main() {
	lambda.Start(Handler)
}
