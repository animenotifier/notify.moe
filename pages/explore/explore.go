package explore

import (
	"strconv"
	"time"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Filter ...
func Filter(ctx *aero.Context) string {
	year := ctx.Get("year")
	season := ctx.Get("season")
	status := ctx.Get("status")
	typ := ctx.Get("type")
	sort := ctx.Get("sort")
	user := utils.GetUser(ctx)
	now := time.Now()

	if year == "" {
		year = strconv.Itoa(now.Year())
	}

	if season == "" {
		season = arn.DateToSeason(now)
	}

	if status == "" {
		status = "current"
	}

	if typ == "" {
		typ = "tv"
	}

	results := filterAnime(year, season, status, typ, sort, user)

	if year == "any" {
		year = ""
	}

	if season == "any" {
		season = ""
	}

	if status == "any" {
		status = ""
	}

	if typ == "any" {
		typ = ""
	}

	if sort == "Popularity" {
		sort = ""
	}

	return ctx.HTML(components.ExploreAnime(results, year, season, status, typ, sort, user))
}

func filterAnime(year, season, status, typ string, sortAlgo string, user *arn.User) []*arn.Anime {
	var results []*arn.Anime

	for anime := range arn.StreamAnime() {
		if year != "any" {
			if len(anime.StartDate) < 4 {
				continue
			}

			if anime.StartDate[:4] != year {
				continue
			}
		}

		if status != "any" && anime.Status != status {
			continue
		}

		if season != "any" && anime.Season() != season {
			continue
		}

		if (typ != "any" || anime.Type == "music" || anime.Type == "tba") && anime.Type != typ {
			continue
		}

		results = append(results, anime)
	}

	arn.SortAnimeWithAlgo(results, status, sortAlgo, user)
	return results
}
