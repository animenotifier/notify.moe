package arn

import "github.com/aerogo/nano"

// ClientErrorReport saves JavaScript errors that happen in web clients like browsers.
type ClientErrorReport struct {
	Message      string `json:"message"`
	Stack        string `json:"stack"`
	FileName     string `json:"fileName"`
	LineNumber   int    `json:"lineNumber"`
	ColumnNumber int    `json:"columnNumber"`

	hasID
	hasCreator
}

// StreamClientErrorReports returns a stream of all client error reports.
func StreamClientErrorReports() <-chan *ClientErrorReport {
	channel := make(chan *ClientErrorReport, nano.ChannelBufferSize)

	go func() {
		for obj := range DB.All("ClientErrorReport") {
			channel <- obj.(*ClientErrorReport)
		}

		close(channel)
	}()

	return channel
}

// AllClientErrorReports returns a slice of all client error reports.
func AllClientErrorReports() []*ClientErrorReport {
	all := make([]*ClientErrorReport, 0, DB.Collection("ClientErrorReport").Count())
	stream := StreamClientErrorReports()

	for obj := range stream {
		all = append(all, obj)
	}

	return all
}
