package explore

import (
	"strconv"
	"time"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
)

// Filter filters the anime for the explore page.
func Filter(ctx aero.Context) error {
	year := ctx.Get("year")
	season := ctx.Get("season")
	status := ctx.Get("status")
	typ := ctx.Get("type")
	user := arn.GetUserFromContext(ctx)
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

	results := filterAnime(year, season, status, typ)

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

	return ctx.HTML(components.ExploreAnime(results, year, season, status, typ, user))
}

func filterAnime(year, season, status, typ string) []*arn.Anime {
	results := make([]*arn.Anime, 0, 1024)

	for anime := range arn.StreamAnime() {
		if anime.IsDraft {
			continue
		}

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

	arn.SortAnimeByQualityDetailed(results, status)
	return results
}
