package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type Event struct {
	source       string
	email        string
	templateName string
}

func Handler(ctx context.Context, event Event) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("sa-east-1"),
	})
	if err != nil {
		panic(err)
	}
	svc := ses.New(sess)

	template, err := getTemplate(event.templateName, map[string]any{
		"name": "Calixto da zs",
	})
	if err != nil {
		panic(err)
	}

	_, err = svc.SendEmail(&ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(event.email),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Data: aws.String(template),
				},
			},
			Subject: &ses.Content{
				Data: aws.String("Testezada :)"),
			},
		},
		Source: aws.String(event.source),
	})
	if err != nil {
		emailNotVerifiedError := strings.Contains(err.Error(), "Email address is not verified")

		if emailNotVerifiedError {
			_, err = svc.VerifyEmailIdentity(&ses.VerifyEmailIdentityInput{
				EmailAddress: aws.String(event.email),
			})
			if err != nil {
				panic(err)
			}
		}
		return err
	}

	return nil
}

func getTemplate(templateName string, data map[string]any) (template string, err error) {
	wd, err := os.Getwd()
	fmt.Println(wd)
	if err != nil {
		return "", err
	}

	templatePath := filepath.Join(wd, "resources", templateName)
	templateBytes, err := os.ReadFile(templatePath)
	if err != nil {
		return "", err
	}

	return string(templateBytes), nil
}

func main() {
	if env, ok := os.LookupEnv("NODE_ENV"); ok && env == "local" {
		fmt.Println(Handler(context.Background(), Event{
			source:       "kauacalixto44@gmail.com",
			email:        "kauacalixtocontato@gmail.com",
			templateName: "teste.html",
		}))

	} else {
		lambda.Start(Handler)
	}
}
