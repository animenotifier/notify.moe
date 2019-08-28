package statistics

import (
	"fmt"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
)

// Anime ...
func Anime(ctx aero.Context) error {
	pieCharts := getAnimeStats()
	return ctx.HTML(components.Statistics(pieCharts))
}

func getAnimeStats() []*arn.PieChart {
	shoboi := stats{}
	anilist := stats{}
	mal := stats{}
	anidb := stats{}
	status := stats{}
	types := stats{}
	rating := stats{}
	twist := stats{}

	for anime := range arn.StreamAnime() {
		if anime.GetMapping("shoboi/anime") != "" {
			shoboi["Connected with Shoboi"]++
		} else {
			shoboi["Not connected with Shoboi"]++
		}

		if anime.GetMapping("anilist/anime") != "" {
			anilist["Connected with AniList"]++
		} else {
			anilist["Not connected with AniList"]++
		}

		if anime.GetMapping("myanimelist/anime") != "" {
			mal["Connected with MyAnimeList"]++
		} else {
			mal["Not connected with MyAnimeList"]++
		}

		if anime.GetMapping("anidb/anime") != "" {
			anidb["Connected with AniDB"]++
		} else {
			anidb["Not connected with AniDB"]++
		}

		rating[fmt.Sprint(int(anime.Rating.Overall+0.5))]++

		found := false
		for _, episode := range anime.Episodes() {
			if episode.Links != nil && episode.Links["twist.moe"] != "" {
				found = true
				break
			}
		}

		if found {
			twist["Connected with AnimeTwist"]++
		} else {
			twist["Not connected with AnimeTwist"]++
		}

		status[anime.Status]++
		types[anime.Type]++
	}

	return []*arn.PieChart{
		arn.NewPieChart("Type", types),
		arn.NewPieChart("Status", status),
		arn.NewPieChart("Rating", rating),
		arn.NewPieChart("MyAnimeList", mal),
		arn.NewPieChart("AniList", anilist),
		arn.NewPieChart("AniDB", anidb),
		arn.NewPieChart("Shoboi", shoboi),
		arn.NewPieChart("AnimeTwist", twist),
	}
}
