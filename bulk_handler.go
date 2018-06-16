package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

// BulkHandler handles incoming request from /track/bulk path
func BulkHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var bulk TrackBulk
	err := json.Unmarshal([]byte(request.Body), &bulk)

	if err != nil {
		return SendError(err)
	}

	response, err := json.Marshal(bulk)
	if err != nil {
		return SendError(err)
	}

	return SendSuccess(string(response))
}
