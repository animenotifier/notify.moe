package main

import (
	"strings"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
)

func main() {
	aero.Parallel(updateAnimeIndex, updateUserIndex)
}

func updateAnimeIndex() {
	animeSearchIndex := arn.NewSearchIndex()

	// Anime
	animeStream, err := arn.AllAnime()

	if err != nil {
		panic(err)
	}

	for anime := range animeStream {
		animeSearchIndex.TextToID[strings.ToLower(anime.Title.Canonical)] = anime.ID
	}

	// Save in database
	err = arn.DB.Set("SearchIndex", "Anime", animeSearchIndex)

	if err != nil {
		panic(err)
	}
}

func updateUserIndex() {
	userSearchIndex := arn.NewSearchIndex()

	// Users
	userStream, err := arn.AllUsers()

	if err != nil {
		panic(err)
	}

	for user := range userStream {
		userSearchIndex.TextToID[strings.ToLower(user.Nick)] = user.ID
	}

	// Save in database
	err = arn.DB.Set("SearchIndex", "User", userSearchIndex)

	if err != nil {
		panic(err)
	}
}
