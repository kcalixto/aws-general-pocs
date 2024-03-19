package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

var test string

func Handler() error {
	time.Sleep(time.Duration(100) * time.Millisecond)
	if len(test) == 0 {
		test += "aloha:"
	} else {
		test += "A"
	}

	fmt.Println(test)
	return nil
}

func main() {
	lambda.Start(Handler)
}
