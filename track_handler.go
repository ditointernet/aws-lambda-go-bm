package main

import (
	"encoding/json"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/uudashr/iso8601"
)

// TrackHandler handles incoming request from /track path
func TrackHandler(request events.APIGatewayProxyRequest) (string, error) {
	var event Track
	err := json.Unmarshal([]byte(request.Body), &event)
	if err != nil {
		return "", err
	}
	event.ReceivedAt = iso8601.Time(time.Now())

	data, err := json.Marshal(event)
	if err != nil {
		return "", err
	}

	client := NewClient()
	response, err := TrackEvent(client, data)

	if err != nil {
		return "", err
	}

	return response.String(), nil
}
