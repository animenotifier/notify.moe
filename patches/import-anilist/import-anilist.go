package main

import (
	"fmt"
	"strings"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

var allAnime []*arn.Anime

func init() {
	allAnime, _ = arn.AllAnime()
}

func main() {
	arn.PanicOnError(arn.AniList.Authorize())
	println(arn.AniList.AccessToken)

	user, _ := arn.GetUserByNick("Boltasar")
	animeList, err := arn.AniList.GetAnimeList(user)
	arn.PanicOnError(err)

	importList(animeList.Lists.Watching)
	importList(animeList.Lists.Completed)
}

func importList(animeListItems []*arn.AniListAnimeListItem) {
	imported := []*arn.Anime{}

	for _, item := range animeListItems {
		anime := findAnimeByName(item.Anime)
		if anime != nil {
			// fmt.Println(item.Anime.TitleRomaji, "=>", anime.Title.Romaji)
			imported = append(imported, anime)
		}
	}

	color.Green("%d / %d", len(imported), len(animeListItems))
}

func findAnimeByName(search *arn.AniListAnime) *arn.Anime {
	var mostSimilar *arn.Anime
	var similarity float64

	for _, anime := range allAnime {
		anime.Title.Japanese = strings.Replace(anime.Title.Japanese, "2ndシーズン", "2", 1)
		anime.Title.Romaji = strings.Replace(anime.Title.Romaji, " 2nd Season", " 2", 1)
		search.TitleJapanese = strings.TrimSpace(strings.Replace(search.TitleJapanese, "2ndシーズン", "2", 1))
		search.TitleRomaji = strings.TrimSpace(strings.Replace(search.TitleRomaji, " 2nd Season", " 2", 1))

		titleSimilarity := arn.StringSimilarity(anime.Title.Romaji, search.TitleRomaji)

		if strings.ToLower(anime.Title.Japanese) == strings.ToLower(search.TitleJapanese) {
			titleSimilarity += 1.0
		}

		if strings.ToLower(anime.Title.Romaji) == strings.ToLower(search.TitleRomaji) {
			titleSimilarity += 1.0
		}

		if strings.ToLower(anime.Title.English) == strings.ToLower(search.TitleEnglish) {
			titleSimilarity += 1.0
		}

		if titleSimilarity > similarity {
			mostSimilar = anime
			similarity = titleSimilarity
		}
	}

	if mostSimilar.EpisodeCount != search.TotalEpisodes {
		similarity -= 0.02
	}

	if similarity >= 0.92 {
		fmt.Printf("MATCH:    %s => %s (%.2f)\n", search.TitleRomaji, mostSimilar.Title.Romaji, similarity)
		return mostSimilar
	}

	color.Red("MISMATCH: %s => %s (%.2f)", search.TitleRomaji, mostSimilar.Title.Romaji, similarity)

	return nil
}
