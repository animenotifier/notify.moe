package filteranime

import (
	"net/http"
	"strings"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
)

const maxAnimeEntries = 70

// editorList renders the anime list with the given title and filter.
func editorList(ctx aero.Context, title string, filter func(*arn.Anime) bool, searchLink func(*arn.Anime) string) error {
	user := arn.GetUserFromContext(ctx)

	if user == nil || (user.Role != "admin" && user.Role != "editor") {
		return ctx.Error(http.StatusUnauthorized, "Not authorized")
	}

	animes, count := filterAnime(ctx, user, filter)

	// Determine URL
	url := strings.TrimPrefix(ctx.Path(), "/_")
	urlParts := strings.Split(url, "/")
	urlParts = urlParts[:len(urlParts)-4]
	url = strings.Join(urlParts, "/")

	return ctx.HTML(components.AnimeEditorListFull(
		title,
		animes,
		count,
		url,
		searchLink,
		user,
	))
}

// filterAnime filters anime by the given filter function and
// additionally applies year and types filters if specified.
func filterAnime(ctx aero.Context, user *arn.User, filter func(*arn.Anime) bool) ([]*arn.Anime, int) {
	year := ctx.Get("year")
	status := ctx.Get("status")
	season := ctx.Get("season")
	typ := ctx.Get("type")

	if year == "any" {
		year = ""
	}

	if status == "any" {
		status = ""
	}

	if season == "any" {
		season = ""
	}

	if typ == "any" {
		typ = ""
	}

	settings := user.Settings()
	settings.Editor.Filter.Year = year
	settings.Editor.Filter.Season = season
	settings.Editor.Filter.Status = status
	settings.Editor.Filter.Type = typ
	settings.Save()

	// Filter
	animes := arn.FilterAnime(func(anime *arn.Anime) bool {
		if year != "" && (len(anime.StartDate) < 4 || anime.StartDate[:4] != year) {
			return false
		}

		if season != "" && anime.Season() != season {
			return false
		}

		if status != "" && anime.Status != status {
			return false
		}

		if typ != "" && anime.Type != typ {
			return false
		}

		return filter(anime)
	})

	// Sort
	arn.SortAnimeByQuality(animes)

	// Limit
	count := len(animes)

	if count > maxAnimeEntries {
		animes = animes[:maxAnimeEntries]
	}

	return animes, count
}
