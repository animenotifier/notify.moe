package arn

import (
	"github.com/aerogo/nano"
	"github.com/animenotifier/kitsu"
)

// StreamKitsuMappings returns a stream of all Kitsu mappings.
func StreamKitsuMappings() <-chan *kitsu.Mapping {
	channel := make(chan *kitsu.Mapping, nano.ChannelBufferSize)

	go func() {
		for obj := range Kitsu.All("Mapping") {
			channel <- obj.(*kitsu.Mapping)
		}

		close(channel)
	}()

	return channel
}

// FilterKitsuMappings filters all Kitsu mappings by a custom function.
func FilterKitsuMappings(filter func(*kitsu.Mapping) bool) []*kitsu.Mapping {
	var filtered []*kitsu.Mapping

	channel := Kitsu.All("Mapping")

	for obj := range channel {
		realObject := obj.(*kitsu.Mapping)

		if filter(realObject) {
			filtered = append(filtered, realObject)
		}
	}

	return filtered
}
