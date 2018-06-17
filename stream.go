package main

import (
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
func TrackEvent(client *kinesis.Kinesis, data []byte) (*kinesis.PutRecordOutput, error) {
	entry := kinesis.PutRecordInput{
		StreamName:   aws.String("lambda-bm"),
		Data:         data,
		PartitionKey: aws.String(uuid.Must(uuid.NewUUID()).String()),
	}

	return client.PutRecord(&entry)
}

// PutRecordsMessage type
type PutRecordsMessage struct {
	Response *kinesis.PutRecordsOutput
	Error    error
}

// TrackBulk sends bulk to Kinesis stream
func TrackBulk(client *kinesis.Kinesis, chunk [][]byte, ch chan<- PutRecordsMessage) {
	var records []*kinesis.PutRecordsRequestEntry
	for _, data := range chunk {
		entry := kinesis.PutRecordsRequestEntry{}
		entry.SetData(data)
		entry.SetPartitionKey(uuid.Must(uuid.NewUUID()).String())

		records = append(records, &entry)
	}

	input := kinesis.PutRecordsInput{
		Records:    records,
		StreamName: aws.String("lambda-bm"),
	}
	response, error := client.PutRecords(&input)

	ch <- PutRecordsMessage{Response: response, Error: error}
}
