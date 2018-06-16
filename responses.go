package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

// SendError sends error response
func SendError(err error) (events.APIGatewayProxyResponse, error) {
  return events.APIGatewayProxyResponse{
    StatusCode: 400,
    Body:       fmt.Sprintf(`{"error": "%s"}`, err.Error()),
    Headers: map[string]string{
      "Content-Type": "application/json",
    },
  }, nil
}

// SendSuccess sends success response
func SendSuccess(response string) (events.APIGatewayProxyResponse, error) {
  return events.APIGatewayProxyResponse{
    StatusCode: 200,
    Body:       response,
    Headers: map[string]string{
      "Content-Type": "application/json",
    },
  }, nil
}
