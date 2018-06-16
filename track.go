package main

import "github.com/uudashr/iso8601"

// Track defines track event structure
type Track struct {
	Action     string       `json:"action"`
	ReceivedAt iso8601.Time `json:"received_at"`
}

// TrackBulk defines a bulk of track events
type TrackBulk struct {
	Records []Track
}
