package profile

import (
	"net/http"
	"strconv"
	"time"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

type stats map[string]float64

// GetStatsByUser shows statistics for a given user.
func GetStatsByUser(ctx *aero.Context) string {
	nick := ctx.Get("nick")
	viewUser, err := arn.GetUserByNick(nick)
	userStats := utils.UserStats{}
	ratings := stats{}
	status := stats{}
	types := stats{}
	years := stats{}

	if err != nil {
		return ctx.Error(http.StatusNotFound, "User not found", err)
	}

	animeList, err := arn.GetAnimeList(viewUser)
	animeList.PrefetchAnime()

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Anime list not found", err)
	}

	for _, item := range animeList.Items {
		currentWatch := item.Episodes * item.Anime().EpisodeLength
		reWatch := item.RewatchCount * item.Anime().EpisodeCount * item.Anime().EpisodeLength
		duration := time.Duration(currentWatch + reWatch)
		userStats.AnimeWatchingTime += duration * time.Minute

		ratings[strconv.Itoa(int(item.Rating.Overall+0.5))]++
		status[item.Status]++
		types[item.Anime().Type]++

		if item.Anime().StartDate != "" {
			year := item.Anime().StartDate[:4]

			if year < "2000" {
				year = "Before 2000"
			}

			years[year]++
		}
	}

	userStats.PieCharts = []*arn.PieChart{
		arn.NewPieChart("Ratings", ratings),
		arn.NewPieChart("Status", status),
		arn.NewPieChart("Types", types),
		arn.NewPieChart("Years", years),
	}

	return ctx.HTML(components.ProfileStats(&userStats, viewUser, utils.GetUser(ctx), ctx.URI()))
}
