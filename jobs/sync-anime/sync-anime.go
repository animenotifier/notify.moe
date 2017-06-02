package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/animenotifier/arn"
	"github.com/animenotifier/kitsu"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Syncing Anime")

	// Get a stream of all anime
	allAnime := kitsu.AllAnime()

	// Iterate over the stream
	for anime := range allAnime {
		sync(anime)
	}
}

func sync(data *kitsu.Anime) {
	anime := arn.Anime{}

	anime.ID, _ = strconv.Atoi(data.ID)
	anime.Type = strings.ToLower(data.Attributes.ShowType)
	anime.Title.Canonical = data.Attributes.CanonicalTitle
	anime.Title.English = data.Attributes.Titles.En
	anime.Title.Japanese = data.Attributes.Titles.JaJp
	anime.Title.Romaji = data.Attributes.Titles.EnJp
	anime.Title.Synonyms = data.Attributes.AbbreviatedTitles
	anime.Image = data.Attributes.PosterImage.Original
	anime.Summary = arn.FixAnimeDescription(data.Attributes.Synopsis)

	if data.Attributes.YoutubeVideoID != "" {
		anime.Trailers = append(anime.Trailers, &arn.AnimeTrailer{
			Service: "Youtube",
			VideoID: data.Attributes.YoutubeVideoID,
		})
	}

	err := anime.Save()

	status := ""

	if err == nil {
		status = color.GreenString("✔")
	} else {
		status = color.RedString("✘")
	}

	fmt.Println(status, anime.ID, anime.Title.Canonical)

}
