package main

import (
	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	// Get a stream of all anime
	allAnime, err := arn.AllAnime()

	if err != nil {
		panic(err)
	}

	// Iterate over the stream
	for _, anime := range allAnime {
		if anime.Mappings == nil {
			anime.Mappings = []*arn.Mapping{}
		}

		if anime.Episodes == nil {
			anime.Episodes = []*arn.AnimeEpisode{}
		}

		err := anime.Save()

		if err != nil {
			color.Red("Error saving anime: %v", err)
		}
	}
}
