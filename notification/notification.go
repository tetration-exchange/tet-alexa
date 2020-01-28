package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-lambda-go/events"
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
		sendSQS(s[0])
	}
}

func main() {
	//a := NewAlexaMessaging("amzn1.application-oa2-client.826a17f54181487e8b6ae926c0f2c544", "664a9b3eac5dca5a0c81f7cc85ef5f3c796ee28720dc6cd6af636a586b9cc5a5")
	lambda.Start(Handler)

	//userID := "amzn1.ask.account.AG5J2M6V6QCUPTB6UJC6D5N7UGOFZNECWDZ6SBFJQVQ6NUMX2UYRAJPMTAZRNX2WUULYL243U3PMIROMAIT6AUPB2PO7VE6WWF2MKUVNDX43TOHIMWQVYGBCJJFXME4QUPSNVOYAFSIPAC2H4OM4QJLN4LOXDEVUGRCM2P65ZXQPSK5HE6XSB5F4ZA7MP72DLDFBZC5DHKGBYZY"
	//a.NewNotification(userID, "hello Tim!")
}
