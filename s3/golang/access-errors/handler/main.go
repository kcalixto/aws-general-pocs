package main

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	mySession := session.Must(session.NewSession())

	// Create a S3 client with additional configuration
	svc := s3.New(mySession, aws.NewConfig().WithRegion("sa-east-1"))

	bucketName := "kcalixto-private-bucket-test"
	fileKey := "print_01"
	obj, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileKey),
	})
	if err != nil {
		keyDoestNotExists := "The specified key does not exist"
		if strings.Contains(err.Error(), keyDoestNotExists) {
			fmt.Println(fmt.Errorf("error while getting object from S3. File %s not found in %s", fileKey, bucketName))
		} else {
			fmt.Println(fmt.Errorf("untreated error: %+v", err.Error()))
		}
	} else {
		fmt.Printf("obj: %+v", obj)
	}
}
