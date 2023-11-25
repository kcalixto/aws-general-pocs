package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type Event struct {
	email string
}

func Handler(ctx context.Context, event Event) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("sa-east-1"),
	})
	if err != nil {
		panic(err)
	}
	svc := ses.New(sess)

	input := &ses.VerifyEmailIdentityInput{
		EmailAddress: aws.String(event.email),
	}

	_, err = svc.VerifyEmailIdentity(input)
	if err != nil {
		panic(err)
	}

	return nil
}

func main() {
	if env, ok := os.LookupEnv("NODE_ENV"); ok && env == "local" {
		fmt.Println(Handler(context.Background(), Event{
			email: "kauacalixtocontato@gmail.com",
		}))

	} else {
		lambda.Start(Handler)
	}
}
