package main

import (
	"fmt"

	"github.com/akyoto/color"
	"github.com/animenotifier/kitsu"
	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	color.Yellow("Importing Kitsu anime")

	defer color.Green("Finished.")
	defer arn.Node.Close()

	// In case we refresh only one anime
	if InvokeShellArgs() {
		return
	}

	// Get a stream of all anime
	allAnime := kitsu.StreamAnime()

	// Iterate over the stream
	for anime := range allAnime {
		sync(anime)
	}
}

func sync(anime *kitsu.Anime) {
	fmt.Println(anime.ID, anime.Attributes.CanonicalTitle)
	arn.Kitsu.Set("Anime", anime.ID, anime)
}

// func importKitsuAnime(data *kitsu.Anime) *arn.Anime {
// 	anime, err := arn.GetAnime(data.ID)

// 	// This stops overwriting existing data
// 	// if anime != nil {
// 	// 	return anime
// 	// }

// 	if err != nil {
// 		if strings.Contains(err.Error(), "not found") {
// 			anime = &arn.Anime{
// 				Title: &arn.MediaTitle{},
// 			}
// 		} else {
// 			panic(err)
// 		}
// 	}

// 	attr := data.Attributes

// 	// General data
// 	anime.ID = data.ID
// 	anime.Type = strings.ToLower(attr.ShowType)
// 	anime.Title.Canonical = attr.CanonicalTitle
// 	anime.Title.English = attr.Titles.En
// 	anime.Title.Romaji = attr.Titles.EnJp
// 	anime.Title.Synonyms = attr.AbbreviatedTitles
// 	anime.StartDate = attr.StartDate
// 	anime.EndDate = attr.EndDate
// 	anime.EpisodeCount = attr.EpisodeCount
// 	anime.EpisodeLength = attr.EpisodeLength
// 	anime.Status = attr.Status

// 	// Status "unreleased" means the same as "upcoming" so we should normalize it
// 	if anime.Status == "unreleased" {
// 		anime.Status = "upcoming"
// 	}

// 	// Normalize image extension to .jpg if .jpeg is used
// 	if anime.Image.Extension == ".jpeg" {
// 		anime.Image.Extension = ".jpg"
// 	}

// 	anime.Summary = arn.FixAnimeDescription(attr.Synopsis)

// 	if anime.Mappings == nil {
// 		anime.Mappings = []*arn.Mapping{}
// 	}

// 	// Prefer Shoboi Japanese titles over Kitsu JP titles
// 	if anime.GetMapping("shoboi/anime") != "" {
// 		// Only take Kitsu title when our JP title is empty
// 		if anime.Title.Japanese == "" {
// 			anime.Title.Japanese = attr.Titles.JaJp
// 		}
// 	} else {
// 		// Update JP title with Kitsu JP title
// 		anime.Title.Japanese = attr.Titles.JaJp
// 	}

// 	// Import mappings
// 	for _, mapping := range data.Mappings {
// 		switch mapping.Attributes.ExternalSite {
// 		case "myanimelist/anime":
// 			anime.SetMapping("myanimelist/anime", mapping.Attributes.ExternalID, "")
// 		case "anidb":
// 			anime.SetMapping("anidb/anime", mapping.Attributes.ExternalID, "")
// 		case "thetvdb", "thetvdb/series":
// 			fmt.Println(mapping.Attributes.ExternalSite, mapping.Attributes.ExternalID)
// 			anime.SetMapping("thetvdb/anime", mapping.Attributes.ExternalID, "")
// 		case "thetvdb/season":
// 			// Ignore
// 		default:
// 			color.Yellow("Unknown mapping: %s %s", mapping.Attributes.ExternalSite, mapping.Attributes.ExternalID)
// 		}
// 	}

// 	return anime

// 	// Download image
// 	response, err := client.Get(attr.PosterImage.Original).End()

// 	if err == nil && response.StatusCode() == http.StatusOK {
// 		anime.SetImageBytes(response.Bytes())
// 	} else {
// 		color.Red("No image for [%s] %s (%d)", anime.ID, anime, response.StatusCode())
// 	}

// 	// Rating
// 	if anime.Rating == nil {
// 		anime.Rating = &arn.AnimeRating{}
// 	}

// 	if anime.Rating.IsNotRated() {
// 		anime.Rating.Reset()
// 	}

// 	// Popularity
// 	if anime.Popularity == nil {
// 		anime.Popularity = &arn.AnimePopularity{}
// 	}

// 	// Trailers
// 	anime.Trailers = []*arn.ExternalMedia{}

// 	if attr.YoutubeVideoID != "" {
// 		anime.Trailers = append(anime.Trailers, &arn.ExternalMedia{
// 			Service:   "Youtube",
// 			ServiceID: attr.YoutubeVideoID,
// 		})
// 	}

// 	// Save in database
// 	anime.Save()

// 	// Episodes
// 	episodes, err := arn.GetAnimeEpisodes(anime.ID)

// 	if err != nil || episodes == nil {
// 		episodes := &arn.AnimeEpisodes{
// 			AnimeID: anime.ID,
// 			Items:   []*arn.Episode{},
// 		}

// 		arn.DB.Set("AnimeEpisodes", anime.ID, episodes)
// 	}

// 	// Relations
// 	relations, _ := arn.GetAnimeRelations(anime.ID)

// 	if relations == nil {
// 		relations := &arn.AnimeRelations{
// 			AnimeID: anime.ID,
// 			Items:   []*arn.AnimeRelation{},
// 		}

// 		arn.DB.Set("AnimeRelations", anime.ID, relations)
// 	}

// 	// Log
// 	fmt.Println(color.GreenString("✔"), anime.ID, anime.Title.Canonical)

// 	return anime
// }
