package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context, event events.APIGatewayProxyRequest)  (error) {
	body := event.Body
	headers := event.Headers

	fmt.Println("fake error")
	fmt.Println("fake ERROR")
	fmt.Println("fake panic")

	if headers["content-type"] == "application/x-www-form-urlencoded" {
		bytes, err := base64.StdEncoding.DecodeString(body)
		if err != nil {
			panic(err)
		}

		// granType := regexp.
		// assertion := wildcard.Match("&assetion=*", string(bytes))

		// fmt.Println(granType)
		// fmt.Println(assertion)

		grantType := "/grant_type=(.*)&/"
		assertion := "/assertion=(.*)/"
		gt := regexp.MustCompile(grantType)
		at := regexp.MustCompile(assertion)
		grant := strings.Split(gt.FindStringSubmatch(string(bytes))[0], "=")[1]
		assertionStr := strings.Split(at.FindStringSubmatch(string(bytes))[0], "=")[1]

		fmt.Println("grant: ", grant)
		fmt.Println("assertion: ", assertionStr)

	} else {
		fmt.Println(headers["content-type"])
	}

	return nil
}

func main() {
	if env, ok := os.LookupEnv("NODE_ENV"); ok && env == "local" {
		fmt.Println(Handler(context.Background(), events.APIGatewayProxyRequest{
			Body: "Z3JhbnRfdHlwZT11cm4lM0FpZXRmJTNBcGFyYW1zJTNBb2F1dGglM0FncmFudC10eXBlJTNBand0LWJlYXJlciZhc3NlcnRpb249ZXlKaGJHY2lPaUpTVXpJMU5pSXNJblI1Y0NJNklrcFhWQ0o5LmV5SmhkV1FpT2lKb2RIUndjem92TDNCeWIzaDVMbUZ3YVM1d2NtVmlZVzVqYnk1amIyMHVZbkl2WVhWMGFDOXpaWEoyWlhJdmRqRXVNUzkwYjJ0bGJpSXNJbk4xWWlJNklqRXpNakkyTnpOa0xUVTVaREl0TkdWbE55MDROMk0zTFdRNE9UaGpOVFJqTmpVd1lpSXNJbWxoZENJNklqRTJOakEyTURFeE56TWlMQ0psZUhBaU9pSXhOall3TmpBME9EYzVJaXdpYW5ScElqb2lNVFkyTURZd01UTTJORGsyTVNJc0luWmxjaUk2SWpFdU1TSjkuVmVFLWpja24yOVpfbGtMU3E5Z1FfVDYxME9zOUN0dTlsb3d2el9nUDVjWnYwYW1ZSVBQSmNsS1BZYktHS2VOYXJPbDd0Nk83ZUFYT2dmRjJfbWFvTVRvNnp4X3FfaWxhQnpDQU94ZTQ4Z2VzWjN3OUtObWZsalBPRVIxYXo5QlIxWmdKVlRpRllUVS1aTmRLQVV5dmV4b3hMaEpLMENKcWx1UjU0c1Vvb2lQYlBDTlljU205M01ORFh3ZjFoQk1KWUczY3p1c3lOYlBrYjVpNm4yT1pidWZObmpUc1M0UHI2RXBUZ3ZWOEFMa3dMbGNtZ2drSm5MS21IZzJEYXZPd0tvcWJTdzFPQ0pmZDZvZVY2OEppc2VwVUNaZ05HQzBmUDlTdGVoX3ZkRFN0OEtoVFVqRTBjZEhMSVltcDBWZ1hyMDJSWmw3VFBIeHNaZXFlS2NjLVhn",
			Headers: map[string]string{
				"accept":            "*/*",
				"accept-encoding":   "gzip, deflate, br",
				"cache-control":     "no-cache",
				"content-length":    "699",
				"content-type":      "application/x-www-form-urlencoded",
				"host":              "4d1fsib816.execute-api.sa-east-1.amazonaws.com",
				"postman-token":     "9531cd44-63f4-4102-ac17-b82ef0bc938f",
				"user-agent":        "PostmanRuntime/7.32.2",
				"x-amzn-trace-id":   "Root=1-645e7c69-4501bbd908cc38e87e69f778",
				"x-forwarded-for":   "189.110.195.146",
				"x-forwarded-port":  "443",
				"x-forwarded-proto": "https",
			},
		}))

	} else {
		lambda.Start(Handler)
	}
}
