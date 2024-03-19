package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context, asyncEvent map[string]interface{}) {
	req := convertoToAPIGatewayProxyRequest(asyncEvent)

	type myType struct {
		Test string `json:"test"`
	}

	var myTypeInstance myType
	err := json.Unmarshal([]byte(req.Body), &myTypeInstance)
	if err != nil {
		fmt.Printf("Error at Handler: %s \n", err.Error())
	}

	// fmt.Printf("asyncEvent: %+v \n", asyncEvent)
	fmt.Printf("test: %s\n", myTypeInstance.Test)
}

func UnmarshalToType[T any](item any) (resItem T, err error) {
	itemJson, err := json.Marshal(item)

	if err != nil {
		return resItem, err
	}

	err = json.Unmarshal(itemJson, &resItem)

	if err != nil {
		return resItem, err
	}

	return resItem, nil
}

func convertoToAPIGatewayProxyRequest(asyncEvent map[string]interface{}) (proxyRequest events.APIGatewayProxyRequest) {
	bodyInterface, ok := asyncEvent["body"]
	if !ok {
		fmt.Printf("Error at convertoToAPIGatewayProxyRequest.body %s \n", "body not found")
	}
	var bodyBytes []byte
	var err error

	switch bodyInterface.(type) {
	case string:
		bodyBytes = []byte(bodyInterface.(string))
	case map[string]interface{}:
		bodyBytes, err = json.Marshal(bodyInterface)
		if err != nil {
			fmt.Printf("Error at convertoToAPIGatewayProxyRequest.body %s \n", err.Error())
		}
	default:
		fmt.Printf("Error at convertoToAPIGatewayProxyRequest.body %s \n", "body type not supported")
	}
	// reqPath, err := utils.UnmarshalToType[map[string]string](asyncEvent["path"])
	// if err != nil {
	// 	log.Error(fmt.Sprintf("Error at convertoToAPIGatewayProxyRequest reqPath: %s", err.Error()))
	// }
	authorizer, err := UnmarshalToType[map[string]any](asyncEvent["enhancedAuthContext"])
	if err != nil {
		fmt.Printf("Error at convertoToAPIGatewayProxyRequest authorized: %s \n", err.Error())
	}
	headers, err := UnmarshalToType[map[string]string](asyncEvent["headers"])
	if err != nil {
		fmt.Printf("Error at convertoToAPIGatewayProxyRequest headers: %s \n", err.Error())
	}

	proxyRequest.Body = string(bodyBytes)
	proxyRequest.RequestContext.Authorizer = authorizer
	// proxyRequest.PathParameters = reqPath
	proxyRequest.Headers = headers
	return proxyRequest
}

func main() {
	lambda.Start(Handler)
}
