package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kcalixto/aws-general-pocs/ses/get-email-template/sesClient"
)

func Handler(ctx context.Context) error {
	c, err := sesClient.NewSESClient(ctx, sesClient.ExternalSESConfig{
		// RoleArn:         "arn:aws:iam::047299492778:role/application-ses-delivery",
		// ExternalID:      "iq-cross-account-key",
		Region:          "sa-east-1",
		// RoleSessionName: "application-ses-delivery",
	})
	if err != nil {
		return err
	}

	templateParsed, err := c.ReadAndParseTemplate(ctx, "cpcartaocaixa-v1-1", map[string]interface{}{
		"first_name": "calixto",
	})
	if err != nil {
		return err
	}

	print(templateParsed)
	return nil
}

func main() {
	if env, ok := os.LookupEnv("NODE_ENV"); ok && env == "local" {
		fmt.Println(Handler(context.Background()))
	} else {
		lambda.Start(Handler)
	}
}
