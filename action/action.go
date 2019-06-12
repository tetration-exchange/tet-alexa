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
	tetration.Secret = "8ab282b6a8fe151c425d0dce37d82917ca54d97a"
	tetration.Key = "80c9efd3bb1749138392d743e29a4a62"
	tetration.Endpoint = "https://vesx-2.insbu.net"
	tetration.Verify = false
	tetration.Prefix = "/openapi/v1"

	return tetration
}

func handleQuarantine(request alexa.Request) alexa.Response {
	tetration := newTetration()
	ip := "172.30.0.52"
	scope := "Default"

	// Create a map of attributes to apply to a workload
	attributes := make(map[string]string)
	attributes["quarantine"] = "alexa"

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

	return alexa.NewSimpleResponse("Quarantine", "Quarantined the device with name aws-chorizo from the network")
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
