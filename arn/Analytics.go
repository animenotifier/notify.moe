package arn

import "github.com/aerogo/nano"

// Analytics stores user-related statistics.
type Analytics struct {
	UserID     string              `json:"userId"`
	General    GeneralAnalytics    `json:"general"`
	Screen     ScreenAnalytics     `json:"screen"`
	System     SystemAnalytics     `json:"system"`
	Connection ConnectionAnalytics `json:"connection"`
}

// GeneralAnalytics stores general information.
type GeneralAnalytics struct {
	TimezoneOffset int `json:"timezoneOffset"`
}

// ScreenAnalytics stores information about the device screen.
type ScreenAnalytics struct {
	Width           int     `json:"width"`
	Height          int     `json:"height"`
	AvailableWidth  int     `json:"availableWidth"`
	AvailableHeight int     `json:"availableHeight"`
	PixelRatio      float64 `json:"pixelRatio"`
}

// SystemAnalytics stores information about the CPU and OS.
type SystemAnalytics struct {
	CPUCount int    `json:"cpuCount"`
	Platform string `json:"platform"`
}

// ConnectionAnalytics stores information about connection speed and ping.
type ConnectionAnalytics struct {
	DownLink      float64 `json:"downLink"`
	RoundTripTime float64 `json:"roundTripTime"`
	EffectiveType string  `json:"effectiveType"`
}

// GetAnalytics returns the analytics for the given user ID.
func GetAnalytics(userID UserID) (*Analytics, error) {
	obj, err := DB.Get("Analytics", userID)

	if err != nil {
		return nil, err
	}

	return obj.(*Analytics), nil
}

// StreamAnalytics returns a stream of all analytics.
func StreamAnalytics() <-chan *Analytics {
	channel := make(chan *Analytics, nano.ChannelBufferSize)

	go func() {
		for obj := range DB.All("Analytics") {
			channel <- obj.(*Analytics)
		}

		close(channel)
	}()

	return channel
}

// AllAnalytics returns a slice of all analytics.
func AllAnalytics() []*Analytics {
	all := make([]*Analytics, 0, DB.Collection("Analytics").Count())

	stream := StreamAnalytics()

	for obj := range stream {
		all = append(all, obj)
	}

	return all
}
