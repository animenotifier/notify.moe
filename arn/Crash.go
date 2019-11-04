package arn

import "github.com/aerogo/nano"

// Crash contains data about server crashes.
type Crash struct {
	Error string `json:"error"`
	Stack string `json:"stack"`
	Path  string `json:"path"`

	hasID
	hasCreator
}

// StreamCrashes returns a stream of all crashes.
func StreamCrashes() <-chan *Crash {
	channel := make(chan *Crash, nano.ChannelBufferSize)

	go func() {
		for obj := range DB.All("Crash") {
			channel <- obj.(*Crash)
		}

		close(channel)
	}()

	return channel
}

// AllCrashes returns a slice of all crashes.
func AllCrashes() []*Crash {
	all := make([]*Crash, 0, DB.Collection("Crash").Count())
	stream := StreamCrashes()

	for obj := range stream {
		all = append(all, obj)
	}

	return all
}
