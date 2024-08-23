package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	dynamodbtypes "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/google/uuid"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Item struct {
	pk      string  `dynamodbav:"pk"`
	sk      string  `dynamodbav:"sk"`
	Persons Persons `dynamodbav:"persons"`
}

func (i Item) MarshalDynamoDBAttributeValue() (dynamodbtypes.AttributeValue, error) {
	i.pk = uuid.New().String()
	i.sk = "sk"

	item, err := attributevalue.MarshalMap(i)
	if err != nil {
		return nil, err
	}

	return &dynamodbtypes.AttributeValueMemberM{Value: item}, nil
}

type Persons []Person

func (i Persons) MarshalDynamoDBAttributeValue() (dynamodbtypes.AttributeValue, error) {
	personsBytes, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(personsBytes))

	return &dynamodbtypes.AttributeValueMemberS{Value: string(personsBytes)}, nil
}

func (i *Persons) UnmarshalDynamoDBAttributeValue(av dynamodbtypes.AttributeValue) error {
	var persons Persons
	err := json.Unmarshal([]byte(av.(*dynamodbtypes.AttributeValueMemberS).Value), &persons)
	if err != nil {
		return err
	}

	*i = persons
	return nil
}

func main() {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion("sa-east-1"),
	)
	if err != nil {
		panic(err.Error())
	}

	client := dynamodb.NewFromConfig(cfg)

	flags := os.Args[1:]

	if flags[0] == "create" {
		create(client)
		fmt.Println("Item created")
	}
	if flags[0] == "get" {
		if len(flags) < 2 {
			panic("You must provide a PK")
		}
		get(client, flags[1])
		fmt.Println("Item retrieved")
	}
}

func get(client *dynamodb.Client, pk string) {
	key := map[string]dynamodbtypes.AttributeValue{
		"pk": &dynamodbtypes.AttributeValueMemberS{Value: pk},
		"sk": &dynamodbtypes.AttributeValueMemberS{Value: "sk"},
	}
	out, err := client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("test-db"),
		Key:       key,
	})
	if err != nil {
		panic(err.Error())
	}

	var itemOut Item
	err = attributevalue.UnmarshalMap(out.Item, &itemOut)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%+v", itemOut)
}

func create(client *dynamodb.Client) {
	rawItem := Item{
		// PK: uuid.New().String(),
		// SK: "sk",
		Persons: []Person{
			{
				Name: "John",
				Age:  30,
			},
			{
				Name: "Jane",
				Age:  25,
			},
			{
				Name: "Doe",
				Age:  40,
			},
		},
	}

	item, err := attributevalue.MarshalMap(rawItem)
	if err != nil {
		panic(err.Error())
	}

	_, err = client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("test-db"),
		Item:      item,
	})
	if err != nil {
		panic(err.Error())
	}
}
