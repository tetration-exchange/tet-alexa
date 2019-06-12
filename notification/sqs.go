package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func sendSQS(msg string) {
	svc := sqs.New(session.New())

	msgInput := &sqs.SendMessageInput{
		MessageBody: aws.String(msg),
		QueueUrl:    aws.String("https://sqs.us-west-2.amazonaws.com/938996165657/tetration-sqs"),
	}
	resp, err := svc.SendMessage(msgInput)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("[Send message] \n%v \n\n", resp)
}
