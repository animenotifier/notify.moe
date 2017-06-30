package main

import (
	"fmt"
	"strings"

	"github.com/aerogo/flow"
	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Updating search index")

	flow.Parallel(updateAnimeIndex, updateUserIndex)

	color.Green("Finished.")
}

func updateAnimeIndex() {
	animeSearchIndex := arn.NewSearchIndex()

	// Anime
	animeStream, err := arn.StreamAnime()

	if err != nil {
		panic(err)
	}

	for anime := range animeStream {
		if anime.Title.Romaji != "" {
			animeSearchIndex.TextToID[strings.ToLower(anime.Title.Romaji)] = anime.ID
		}

		if anime.Title.English != "" {
			animeSearchIndex.TextToID[strings.ToLower(anime.Title.English)] = anime.ID
		}

		if anime.Title.Japanese != "" {
			animeSearchIndex.TextToID[strings.ToLower(anime.Title.Japanese)] = anime.ID
		}

		for _, synonym := range anime.Title.Synonyms {
			synonym = strings.ToLower(synonym)

			if synonym != "" && len(synonym) <= 10 {
				animeSearchIndex.TextToID[synonym] = anime.ID
			}
		}
	}

	fmt.Println(len(animeSearchIndex.TextToID), "anime titles")

	// Save in database
	err = arn.DB.Set("SearchIndex", "Anime", animeSearchIndex)

	if err != nil {
		panic(err)
	}
}

func updateUserIndex() {
	userSearchIndex := arn.NewSearchIndex()

	// Users
	userStream, err := arn.StreamUsers()

	if err != nil {
		panic(err)
	}

	for user := range userStream {
		if user.IsActive() && user.Nick != "" {
			userSearchIndex.TextToID[strings.ToLower(user.Nick)] = user.ID
		}
	}

	fmt.Println(len(userSearchIndex.TextToID), "user names")

	// Save in database
	err = arn.DB.Set("SearchIndex", "User", userSearchIndex)

	if err != nil {
		panic(err)
	}
}
