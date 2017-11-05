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
	defer arn.Node.Close()

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
	for anime := range arn.StreamAnime() {
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
	arn.DB.Set("SearchIndex", "Anime", animeSearchIndex)
}

func updateUserIndex() {
	userSearchIndex := arn.NewSearchIndex()

	// Users
	for user := range arn.StreamUsers() {
		if user.HasNick() {
			userSearchIndex.TextToID[strings.ToLower(user.Nick)] = user.ID
		}
	}

	fmt.Println(len(userSearchIndex.TextToID), "user names")

	// Save in database
	arn.DB.Set("SearchIndex", "User", userSearchIndex)
}

func updatePostIndex() {
	postSearchIndex := arn.NewSearchIndex()

	// Users
	for post := range arn.StreamPosts() {
		postSearchIndex.TextToID[strings.ToLower(post.Text)] = post.ID
	}

	fmt.Println(len(postSearchIndex.TextToID), "posts")

	// Save in database
	arn.DB.Set("SearchIndex", "Post", postSearchIndex)
}

func updateThreadIndex() {
	threadSearchIndex := arn.NewSearchIndex()

	// Users
	for thread := range arn.StreamThreads() {
		threadSearchIndex.TextToID[strings.ToLower(thread.Title)] = thread.ID
		threadSearchIndex.TextToID[strings.ToLower(thread.Text)] = thread.ID
	}

	fmt.Println(len(threadSearchIndex.TextToID)/2, "threads")

	// Save in database
	arn.DB.Set("SearchIndex", "Thread", threadSearchIndex)
}
