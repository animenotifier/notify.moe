package main

import (
	"fmt"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

var userName = "Akyoto"
var allAnime []*arn.Anime

func init() {
	allAnime, _ = arn.AllAnime()
}

func main() {
	arn.PanicOnError(arn.AniList.Authorize())
	println(arn.AniList.AccessToken)

	user, _ := arn.GetUserByNick(userName)
	animeList, err := arn.AniList.GetAnimeList(user)
	arn.PanicOnError(err)

	importList(animeList.Lists.Watching)
	importList(animeList.Lists.Completed)
}

func importList(animeListItems []*arn.AniListAnimeListItem) {
	imported := []*arn.Anime{}

	for _, item := range animeListItems {
		anime := arn.FindAniListAnime(item.Anime, allAnime)

		if anime != nil {
			fmt.Println(item.Anime.TitleRomaji, "=>", anime.Title.Romaji)
			imported = append(imported, anime)
		}
	}

	color.Green("%d / %d", len(imported), len(animeListItems))
}
