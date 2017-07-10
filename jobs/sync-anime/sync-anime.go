package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/animenotifier/arn"
	"github.com/animenotifier/kitsu"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Syncing Anime")

	// In case we refresh only one anime
	if InvokeShellArgs() {
		color.Green("Finished.")
		return
	}

	// Get a stream of all anime
	allAnime := kitsu.StreamAnimeWithMappings()

	// Iterate over the stream
	for anime := range allAnime {
		sync(anime)
	}

	color.Green("Finished.")
}

func sync(data *kitsu.Anime) *arn.Anime {
	anime, err := arn.GetAnime(data.ID)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			anime = &arn.Anime{
				Title: &arn.AnimeTitle{},
				Image: &arn.AnimeImageTypes{},
			}
		} else {
			panic(err)
		}
	}

	attr := data.Attributes

	// General data
	anime.ID = data.ID
	anime.Type = strings.ToLower(attr.ShowType)
	anime.Title.Canonical = attr.CanonicalTitle
	anime.Title.English = attr.Titles.En
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
	anime.Summary = arn.FixAnimeDescription(attr.Synopsis)

	if anime.Mappings == nil {
		anime.Mappings = []*arn.Mapping{}
	}

	// Prefer Shoboi Japanese titles over Kitsu JP titles
	if anime.GetMapping("shoboi/anime") != "" {
		// Only take Kitsu title when our JP title is empty
		if anime.Title.Japanese == "" {
			anime.Title.Japanese = attr.Titles.JaJp
		}
	} else {
		// Update JP title with Kitsu JP title
		anime.Title.Japanese = attr.Titles.JaJp
	}

	// Import mappings
	for _, mapping := range data.Mappings {
		switch mapping.Attributes.ExternalSite {
		case "myanimelist/anime":
			anime.AddMapping("myanimelist/anime", mapping.Attributes.ExternalID, "")
		case "anidb":
			anime.AddMapping("anidb/anime", mapping.Attributes.ExternalID, "")
		case "thetvdb/series":
			anime.AddMapping("thetvdb/anime", mapping.Attributes.ExternalID, "")
		case "thetvdb/season":
			// Ignore
		default:
			color.Yellow("Unknown mapping: %s %s", mapping.Attributes.ExternalSite, mapping.Attributes.ExternalID)
		}
	}

	// NSFW
	if attr.Nsfw {
		anime.NSFW = 1
	} else {
		anime.NSFW = 0
	}

	// Rating
	if anime.Rating == nil {
		anime.Rating = &arn.AnimeRating{}
	}

	if anime.Rating.IsNotRated() {
		anime.Rating.Reset()
	}

	// Trailers
	anime.Trailers = []*arn.ExternalMedia{}

	if attr.YoutubeVideoID != "" {
		anime.Trailers = append(anime.Trailers, &arn.ExternalMedia{
			Service:   "Youtube",
			ServiceID: attr.YoutubeVideoID,
		})
	}

	// Save in database
	err = anime.Save()
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

	return anime
}
