package statistics

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Anime ...
func Anime(ctx *aero.Context) string {
	allAnime, err := arn.AllAnime()

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Couldn't fetch anime", err)
	}

	shoboi := stats{}
	anilist := stats{}
	status := stats{}
	types := stats{}
	shoboiEdits := stats{}
	anilistEdits := stats{}
	rating := stats{}

	for _, anime := range allAnime {
		for _, external := range anime.Mappings {
			if external.Service == "shoboi/anime" {
				if external.CreatedBy == "" {
					shoboiEdits["Bot"]++
				} else {
					user, _ := arn.GetUser(external.CreatedBy)
					shoboiEdits[user.Nick]++
				}
			}

			if external.Service == "anilist/anime" {
				if external.CreatedBy == "" {
					anilistEdits["Bot"]++
				} else {
					user, _ := arn.GetUser(external.CreatedBy)
					anilistEdits[user.Nick]++
				}
			}
		}

		if anime.GetMapping("shoboi/anime") != "" {
			shoboi["Connected with Shoboi"]++
		} else {
			shoboi["Not connected with Shoboi"]++
		}

		if anime.GetMapping("anilist/anime") != "" {
			anilist["Connected with Anilist"]++
		} else {
			anilist["Not connected with Anilist"]++
		}

		rating[arn.ToString(int(anime.Rating.Overall+0.5))]++

		status[anime.Status]++
		types[anime.Type]++
	}

	return ctx.HTML(components.Statistics(
		utils.NewPieChart("Type", types),
		utils.NewPieChart("Status", status),
		utils.NewPieChart("Rating", rating),
		utils.NewPieChart("Anilist", anilist),
		utils.NewPieChart("Shoboi", shoboi),
		utils.NewPieChart("Anilist Editors", anilistEdits),
		utils.NewPieChart("Shoboi Editors", shoboiEdits),
	))
}
