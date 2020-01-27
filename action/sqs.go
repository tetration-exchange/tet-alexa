package main

import (
	"fmt"

	alexa "github.com/arienmalec/alexa-go"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func handleNotification(request alexa.Request) alexa.Response {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sqs.New(sess)
	qURL := "https://sqs.us-west-2.amazonaws.com/938996165657/tetration-sqs"

	result, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            &qURL,
		MaxNumberOfMessages: aws.Int64(5),
		VisibilityTimeout:   aws.Int64(20), // 20 seconds
		WaitTimeSeconds:     aws.Int64(0),
	})

	if err != nil {
		fmt.Println("Error", err)
		return alexa.NewSimpleResponse("Couldn't read", "I'm sorry, I could not read notifications")
	}

	if len(result.Messages) == 0 {
		return alexa.NewSimpleResponse("No messages", "There are no notifications")
	}

	var sensorName string

	for _, m := range result.Messages {
		// *string to string
		sensorName = *m.Body
		return alexa.NewSimpleResponse("Notification received",
			fmt.Sprintf(
				"There are notifications available, the first alert is that a workload with hostname %s has tampered with I P tables, we recommend you quarantine the workload.",
				sensorName))
	}
	return alexa.NewSimpleResponse("No messages", "There are no notifications")
}
