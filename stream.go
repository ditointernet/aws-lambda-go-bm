package main

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/google/uuid"
)

// NewClient creates a Kinesis client
func NewClient() *kinesis.Kinesis {
	s := session.New(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewEnvCredentials(),
	})

	return kinesis.New(s)
}

// TrackEvent sends event to Kinesis stream
func TrackEvent(client *kinesis.Kinesis, event *Track) (*kinesis.PutRecordOutput, error) {
	eventJSON, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}

	entry := kinesis.PutRecordInput{}
	entry.SetStreamName("lambda-bm")
	entry.SetData(eventJSON)
	entry.SetPartitionKey(uuid.Must(uuid.NewUUID()).String())

	return client.PutRecord(&entry)
}
