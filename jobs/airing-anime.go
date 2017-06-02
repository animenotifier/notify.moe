// package main

// import (
// 	"fmt"
// 	"sort"

// 	"github.com/animenotifier/arn"
// 	"github.com/fatih/color"
// )

// // AiringAnime ...
// func AiringAnime() {
// 	fmt.Println("Running background job: Airing Anime")

// 	animeList, err := arn.GetAiringAnime()

// 	if err != nil {
// 		color.Red("Failed fetching airing anime")
// 		color.Red(err.Error())
// 		return
// 	}

// 	sort.Sort(arn.AnimeByPopularity(animeList))

// 	// Convert to small anime list
// 	var animeListSmall []*arn.AnimeSmall

// 	for _, anime := range animeList {
// 		animeListSmall = append(animeListSmall, &arn.AnimeSmall{
// 			ID:       anime.ID,
// 			Title:    anime.Title,
// 			Image:    anime.Image,
// 			Watching: anime.Watching,
// 		})
// 	}

// 	saveErr := arn.SetObject("Cache", "airingAnime", &arn.AiringAnimeCacheSmall{
// 		Anime: animeListSmall,
// 	})

// 	if saveErr != nil {
// 		color.Red("Error saving airing anime")
// 		color.Red(saveErr.Error())
// 		return
// 	}
// }
