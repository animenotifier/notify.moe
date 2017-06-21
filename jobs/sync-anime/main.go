package main

import (
	"encoding/json"
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

	println("Finished.")
}

func sync(data *kitsu.Anime) {
	anime := arn.Anime{}
	attr := data.Attributes

	// General data
	anime.ID = data.ID
	anime.Type = strings.ToLower(attr.ShowType)
	anime.Title.Canonical = attr.CanonicalTitle
	anime.Title.English = attr.Titles.En
	anime.Title.Japanese = attr.Titles.JaJp
	anime.Title.Romaji = attr.Titles.EnJp
	anime.Title.Synonyms = attr.AbbreviatedTitles
	anime.Image.Tiny = kitsu.FixImageURL(attr.PosterImage.Tiny)
	anime.Image.Small = kitsu.FixImageURL(attr.PosterImage.Small)
	anime.Image.Large = kitsu.FixImageURL(attr.PosterImage.Large)
	anime.Image.Original = kitsu.FixImageURL(attr.PosterImage.Original)
	anime.StartDate = attr.StartDate
	anime.EndDate = attr.EndDate
	anime.EpisodeCount = attr.EpisodeCount
	anime.EpisodeLength = attr.EpisodeLength
	anime.Status = attr.Status
	anime.NSFW = attr.Nsfw
	anime.Summary = arn.FixAnimeDescription(attr.Synopsis)

	// Rating
	overall, convertError := strconv.ParseFloat(attr.AverageRating, 64)

	if convertError != nil {
		overall = 0
	}

	anime.Rating.Overall = overall

	// Trailers
	anime.Trailers = []arn.AnimeTrailer{}

	if attr.YoutubeVideoID != "" {
		anime.Trailers = append(anime.Trailers, arn.AnimeTrailer{
			Service: "Youtube",
			VideoID: attr.YoutubeVideoID,
		})
	}

	// Save in database
	err := anime.Save()
	status := ""

	if err == nil {
		status = color.GreenString("✔")
	} else {
		color.Red(err.Error())

		data, _ := json.MarshalIndent(anime, "", "\t")
		fmt.Println(string(data))

		status = color.RedString("✘")
	}

	// Log
	fmt.Println(status, anime.ID, anime.Title.Canonical)
}
