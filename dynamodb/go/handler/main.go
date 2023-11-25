package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func Handler(ctx context.Context) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("sa-east-1"),
	})
	if err != nil {
		panic(err)
	}

	return nil
}

func main() {
	if env, ok := os.LookupEnv("NODE_ENV"); ok && env == "local" {
		fmt.Println(Handler(context.Background()))
	} else {
		lambda.Start(Handler)
	}
}
