package main

import (
	alexa "github.com/arienmalec/alexa-go"
	"github.com/aws/aws-lambda-go/lambda"
)

// DispatchIntents dispatches each intent to the right handler
func DispatchIntents(request alexa.Request) alexa.Response {
	var response alexa.Response
	switch request.Body.Intent.Name {
	case "EscapedFlowsIntent":
		response = handleEscapedFlows(request)
	case "QuarantineIntent":
		response = handleQuarantine(request)
	case alexa.HelpIntent:
		response = handleHelp()
	default:
		response = alexa.NewSimpleResponse("Unknown command", "I can't ask Tetration to do that, yet")
	}

	return response
}

func handleEscapedFlows(request alexa.Request) alexa.Response {
	return alexa.NewSimpleResponse("Escaped Flows", "Escaped Flows are Unauthorized Traffic on the network. BOOOM.")
}

func handleQuarantine(request alexa.Request) alexa.Response {
	return alexa.NewSimpleResponse("Quarantine", "Quarantined the device from the network")
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
