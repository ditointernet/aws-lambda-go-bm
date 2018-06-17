package main

import (
	"encoding/json"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/uudashr/iso8601"
)

// BulkHandler handles incoming request from /track/bulk path
func BulkHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var bulk Bulk
	err := json.Unmarshal([]byte(request.Body), &bulk)

	if err != nil {
		return SendError(err)
	}

	chunks, err := generateChunks(bulk.Records, 1)
	if err != nil {
		return SendError(err)
	}

	client := NewClient()
	ch := make(chan PutRecordsMessage)

	for _, chunk := range chunks {
		go TrackBulk(client, chunk, ch)
	}

	var failedRecordCount int64
	var recordsResult []*kinesis.PutRecordsResultEntry

	for i := 0; i < len(chunks); i++ {
		message := <-ch
		failedRecordCount += *message.Response.FailedRecordCount
		recordsResult = append(recordsResult, message.Response.Records...)
	}

	response := kinesis.PutRecordsOutput{
		FailedRecordCount: &failedRecordCount,
		Records:           recordsResult,
	}
	return SendSuccess(response.String())
}

func generateChunks(records []Track, chunkSize int) ([][][]byte, error) {
	var chunks [][][]byte
	var chunk [][]byte

	now := iso8601.Time(time.Now())
	for _, record := range records {
		record.ReceivedAt = now
		data, err := json.Marshal(record)
		if err != nil {
			return nil, err
		}

		if len(chunk) >= chunkSize {
			chunks = append(chunks, chunk)
			chunk = nil
		}

		chunk = append(chunk, data)
	}

	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}

	return chunks, nil
}
