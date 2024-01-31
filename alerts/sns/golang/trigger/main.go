package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context, event events.APIGatewayProxyRequest) error {
	return fmt.Errorf("fake-error")
}

func main() {
	lambda.Start(Handler)
}
