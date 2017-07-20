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

	flow.Parallel(
		updateAnimeIndex,
		updateUserIndex,
		updatePostIndex,
		updateThreadIndex,
	)

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
		if anime.Title.Canonical != "" {
			animeSearchIndex.TextToID[strings.ToLower(anime.Title.Canonical)] = anime.ID
		}

		if anime.Title.Romaji != "" {
			animeSearchIndex.TextToID[strings.ToLower(anime.Title.Romaji)] = anime.ID
		}

		// Make sure we only include Japanese titles that
		// don't overlap with the English titles.
		if anime.Title.Japanese != "" && animeSearchIndex.TextToID[strings.ToLower(anime.Title.Japanese)] == "" {
			animeSearchIndex.TextToID[strings.ToLower(anime.Title.Japanese)] = anime.ID
		}

		// Same with English titles, don't overwrite other stuff.
		if anime.Title.English != "" && animeSearchIndex.TextToID[strings.ToLower(anime.Title.English)] == "" {
			animeSearchIndex.TextToID[strings.ToLower(anime.Title.English)] = anime.ID
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
	arn.PanicOnError(err)

	for user := range userStream {
		if user.HasNick() {
			userSearchIndex.TextToID[strings.ToLower(user.Nick)] = user.ID
		}
	}

	fmt.Println(len(userSearchIndex.TextToID), "user names")

	// Save in database
	err = arn.DB.Set("SearchIndex", "User", userSearchIndex)
	arn.PanicOnError(err)
}

func updatePostIndex() {
	postSearchIndex := arn.NewSearchIndex()

	// Users
	postStream, err := arn.StreamPosts()
	arn.PanicOnError(err)

	for post := range postStream {
		postSearchIndex.TextToID[strings.ToLower(post.Text)] = post.ID
	}

	fmt.Println(len(postSearchIndex.TextToID), "posts")

	// Save in database
	err = arn.DB.Set("SearchIndex", "Post", postSearchIndex)
	arn.PanicOnError(err)
}

func updateThreadIndex() {
	threadSearchIndex := arn.NewSearchIndex()

	// Users
	threadStream, err := arn.StreamThreads()
	arn.PanicOnError(err)

	for thread := range threadStream {
		threadSearchIndex.TextToID[strings.ToLower(thread.Title)] = thread.ID
		threadSearchIndex.TextToID[strings.ToLower(thread.Text)] = thread.ID
	}

	fmt.Println(len(threadSearchIndex.TextToID)/2, "threads")

	// Save in database
	err = arn.DB.Set("SearchIndex", "Thread", threadSearchIndex)
	arn.PanicOnError(err)
}
