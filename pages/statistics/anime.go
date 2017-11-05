package statistics

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

// Anime ...
func Anime(ctx *aero.Context) string {
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
	shoboiEdits := stats{}
	anilistEdits := stats{}
	malEdits := stats{}
	anidbEdits := stats{}
	rating := stats{}
	twist := stats{}

	for anime := range arn.StreamAnime() {
		for _, external := range anime.Mappings {
			if external.Service == "shoboi/anime" {
				if external.CreatedBy == "" {
					shoboiEdits["(auto-generated)"]++
				} else {
					user, err := arn.GetUser(external.CreatedBy)
					arn.PanicOnError(err)
					shoboiEdits[user.Nick]++
				}
			}

			if external.Service == "anilist/anime" {
				if external.CreatedBy == "" {
					anilistEdits["(auto-generated)"]++
				} else {
					user, err := arn.GetUser(external.CreatedBy)
					arn.PanicOnError(err)
					anilistEdits[user.Nick]++
				}
			}

			if external.Service == "myanimelist/anime" {
				if external.CreatedBy == "" {
					malEdits["(auto-generated)"]++
				} else {
					user, err := arn.GetUser(external.CreatedBy)
					arn.PanicOnError(err)
					malEdits[user.Nick]++
				}
			}

			if external.Service == "anidb/anime" {
				if external.CreatedBy == "" {
					anidbEdits["(auto-generated)"]++
				} else {
					user, err := arn.GetUser(external.CreatedBy)
					arn.PanicOnError(err)
					anidbEdits[user.Nick]++
				}
			}
		}

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

		rating[arn.ToString(int(anime.Rating.Overall+0.5))]++

		found := false
		for _, episode := range anime.Episodes().Items {
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
		// arn.NewPieChart("MyAnimeList Editors", malEdits),
		arn.NewPieChart("AniList Editors", anilistEdits),
		// arn.NewPieChart("AniDB Editors", anidbEdits),
		arn.NewPieChart("Shoboi Editors", shoboiEdits),
	}
}
