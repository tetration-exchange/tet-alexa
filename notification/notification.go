package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler is the lambda hander
func Handler(kinesisEvent events.KinesisEvent) {
	for _, record := range kinesisEvent.Records {
		kinesisRecord := record.Kinesis
		dataBytes := kinesisRecord.Data
		dataText := string(dataBytes)

		fmt.Printf("%s Data = %s \n", record.EventName, dataText)

		var jMap map[string]interface{}
		if err := json.Unmarshal(dataBytes, &jMap); err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}

		//alertText := jMap["alertText"]
		s := strings.Fields(jMap["alertText"].(string))
		sendSNS(s[0])
	}
}

func main() {
	lambda.Start(Handler)
}
