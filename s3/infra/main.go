package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, event events.S3Event) error {
	for _, record := range event.Records {
		fmt.Printf("record: %v\n", record)
	}

	return nil
}

func main() {
	lambda.Start(handler)
}
