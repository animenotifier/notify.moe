package arn

import (
	"fmt"

	"github.com/aerogo/nano"
)

// IgnoreAnimeDifferenceEditorScore represents how many points you get for a diff ignore.
const IgnoreAnimeDifferenceEditorScore = 2

// IgnoreAnimeDifference saves which differences between anime databases can be ignored.
type IgnoreAnimeDifference struct {
	// The ID is built like this: arn:323|mal:356|JapaneseTitle
	ID        string `json:"id"`
	ValueHash uint64 `json:"valueHash"`

	hasCreator
}

// GetIgnoreAnimeDifference ...
func GetIgnoreAnimeDifference(id string) (*IgnoreAnimeDifference, error) {
	obj, err := DB.Get("IgnoreAnimeDifference", id)

	if err != nil {
		return nil, err
	}

	return obj.(*IgnoreAnimeDifference), nil
}

// CreateDifferenceID ...
func CreateDifferenceID(animeID string, dataProvider string, malAnimeID string, typeName string) string {
	return fmt.Sprintf("arn:%s|%s:%s|%s", animeID, dataProvider, malAnimeID, typeName)
}

// IsAnimeDifferenceIgnored tells you whether the given difference is being ignored.
func IsAnimeDifferenceIgnored(animeID string, dataProvider string, malAnimeID string, typeName string, hash uint64) bool {
	key := CreateDifferenceID(animeID, dataProvider, malAnimeID, typeName)
	ignore, err := GetIgnoreAnimeDifference(key)

	if err != nil {
		return false
	}

	return ignore.ValueHash == hash
}

// StreamIgnoreAnimeDifferences returns a stream of all ignored differences.
func StreamIgnoreAnimeDifferences() <-chan *IgnoreAnimeDifference {
	channel := make(chan *IgnoreAnimeDifference, nano.ChannelBufferSize)

	go func() {
		for obj := range DB.All("IgnoreAnimeDifference") {
			channel <- obj.(*IgnoreAnimeDifference)
		}

		close(channel)
	}()

	return channel
}

// AllIgnoreAnimeDifferences returns a slice of all ignored differences.
func AllIgnoreAnimeDifferences() []*IgnoreAnimeDifference {
	all := make([]*IgnoreAnimeDifference, 0, DB.Collection("IgnoreAnimeDifference").Count())

	stream := StreamIgnoreAnimeDifferences()

	for obj := range stream {
		all = append(all, obj)
	}

	return all
}

// FilterIgnoreAnimeDifferences filters all ignored differences by a custom function.
func FilterIgnoreAnimeDifferences(filter func(*IgnoreAnimeDifference) bool) []*IgnoreAnimeDifference {
	var filtered []*IgnoreAnimeDifference

	for obj := range StreamIgnoreAnimeDifferences() {
		if filter(obj) {
			filtered = append(filtered, obj)
		}
	}

	return filtered
}
