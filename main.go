package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var routes = map[string]func(events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error){
	"/track":      TrackHandler,
	"/track/bulk": BulkHandler,
}

func router(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return routes[request.Path](request)
}

func main() {
	lambda.Start(router)
}
