package main

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context) error {
	return random(func() error { return fmt.Errorf("error") })
}

func main() {
	lambda.Start(Handler)
}

func random(fn func() error) error {
	chance := 66
	if rand.Intn(100) < chance {
		return fn()
	}

	return nil
}
