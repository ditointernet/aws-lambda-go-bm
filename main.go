package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var routes = map[string]func(events.APIGatewayProxyRequest) (string, error){
	"/track":      TrackHandler,
	"/track/bulk": BulkHandler,
}

func router(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	response, err := routes[request.Path](request)
	if err != nil {
		return SendError(err)
	}

	return SendSuccess(response)
}

func main() {
	lambda.Start(router)
}
