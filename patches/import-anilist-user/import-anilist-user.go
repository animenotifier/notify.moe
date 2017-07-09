package main

import (
	"fmt"

	"github.com/animenotifier/anilist"
	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

var userName = "Akyoto"
var allAnime []*arn.Anime

func init() {
	allAnime, _ = arn.AllAnime()
}

func main() {
	arn.PanicOnError(anilist.Authorize())
	println(anilist.AccessToken)

	user, _ := arn.GetUserByNick(userName)
	animeList, err := anilist.GetAnimeList(user.Accounts.AniList.Nick)
	arn.PanicOnError(err)

	importList(animeList.Lists.Watching)
	importList(animeList.Lists.Completed)
}

func importList(animeListItems []*anilist.AnimeListItem) {
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
