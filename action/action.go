package main

import (
	alexa "github.com/arienmalec/alexa-go"
	"github.com/aws/aws-lambda-go/lambda"
	goh4 "github.com/remiphilippe/go-h4"
)

// DispatchIntents dispatches each intent to the right handler
func DispatchIntents(request alexa.Request) alexa.Response {
	var response alexa.Response
	switch request.Body.Intent.Name {
	case "NotificationIntent":
		response = handleNotification(request)
	case "QuarantineIntent":
		response = handleQuarantine(request)
	case alexa.HelpIntent:
		response = handleHelp()
	default:
		response = alexa.NewSimpleResponse("Unknown command", "I can't ask Tetration to do that, yet")
	}

	return response
}

func newTetration() *goh4.H4 {

	tetration := new(goh4.H4)
	tetration.Secret = "17e49e0f84a1acce67998c33a61a298fb8dfd98c"
	tetration.Key = "b21a1efff5b64e79a970222b15e7bc4c"
	tetration.Endpoint = "https://demo.tetrationpreview.com"
	tetration.Verify = false
	tetration.Prefix = "/openapi/v1"

	return tetration
}

func handleQuarantine(request alexa.Request) alexa.Response {
	tetration := newTetration()
	ip := "10.66.2.83"
	scope := "demo"

	// Create a map of attributes to apply to a workload
	attributes := make(map[string]string)
	attributes["quarantine"] = "true"

	// Define where we want to attach this annotation
	annotation := goh4.Annotation{
		IP:         ip,
		Attributes: attributes,
		Scope:      scope,
	}

	err := tetration.AddAnnotation(&annotation)
	if err != nil {
		return alexa.NewSimpleResponse("Failed to quarantine", "I could not quarantine the device from the network")
	}

	return alexa.NewSimpleResponse("Quarantine", "Quarantined the suspicious device from the network")
}

func handleHelp() alexa.Response {
	return alexa.NewSimpleResponse("Help for Tetration", "Tetration Alexa App")
}

// Handler is the lambda hander
func Handler(request alexa.Request) (alexa.Response, error) {
	return DispatchIntents(request), nil
}

func main() {
	lambda.Start(Handler)
}
