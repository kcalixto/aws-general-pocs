package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
)

func Handler(ctx context.Context) error {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion("sa-east-1"),
	)
	if err != nil {
		return err
	}

	client := dynamodb.NewFromConfig(cfg)

	item, err := attributevalue.MarshalMap(map[string]any{
		"pk":         "item",
		"created_at": time.Now().Unix(),
	})
	if err != nil {
		return err
	}

	_, err = client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
		Item:      item,
	})
	if err != nil {
		return err
	}

	key := map[string]types.AttributeValue{
		"pk": &types.AttributeValueMemberS{Value: "item"},
	}
	out, err := client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
		Key:       key,
	})
	if err != nil {
		return err
	}

	var itemOut map[string]any
	err = attributevalue.UnmarshalMap(out.Item, &itemOut)
	if err != nil {
		return err
	}

	fmt.Printf("item: %v\n", itemOut["value"])
	fmt.Printf("item.created_at type: %T\n", itemOut["created_at"])

	return nil
}

func main() {
	if env, ok := os.LookupEnv("NODE_ENV"); ok && env == "local" {
		fmt.Println(Handler(context.Background()))
	} else {
		lambda.Start(Handler)
	}
}
