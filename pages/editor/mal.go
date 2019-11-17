package editor

import (
	"sort"

	"github.com/animenotifier/notify.moe/utils/animediff"

	"github.com/aerogo/aero"
	"github.com/animenotifier/mal"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

const maxCompareMALEntries = 15

// diffFunction is the signature of a diff function.
type diffFunction func(*arn.Anime, *mal.Anime) []animediff.Difference

// CompareMAL ...
func CompareMAL(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)
	year := ctx.Get("year")
	status := ctx.Get("status")
	season := ctx.Get("season")
	typ := ctx.Get("type")

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

	settings := user.Settings()
	settings.Editor.Filter.Year = year
	settings.Editor.Filter.Season = season
	settings.Editor.Filter.Status = status
	settings.Editor.Filter.Type = typ
	settings.Save()

	animes := arn.FilterAnime(func(anime *arn.Anime) bool {
		if year != "" && (len(anime.StartDate) < 4 || anime.StartDate[:4] != year) {
			return false
		}

		if status != "" && anime.Status != status {
			return false
		}

		if season != "" && anime.Season() != season {
			return false
		}

		if typ != "" && anime.Type != typ {
			return false
		}

		return anime.GetMapping("myanimelist/anime") != ""
	})

	sort.Slice(animes, func(i, j int) bool {
		a := animes[i]
		b := animes[j]

		aPop := a.Popularity.Total()
		bPop := b.Popularity.Total()

		if aPop == bPop {
			return a.Title.Canonical < b.Title.Canonical
		}

		return aPop > bPop
	})

	comparisons := compare(animes)

	return ctx.HTML(components.CompareMAL(comparisons, year, status, typ, "/editor/mal/diff/anime", user))
}

// compare builds the comparisons to MAL anime entries.
func compare(animes []*arn.Anime) []*utils.MALComparison {
	comparisons := []*utils.MALComparison{}
	malAnimeCollection := arn.MAL.Collection("Anime")

	for _, anime := range animes {
		malID := anime.GetMapping("myanimelist/anime")
		obj, err := malAnimeCollection.Get(malID)

		if err != nil {
			continue
		}

		malAnime := obj.(*mal.Anime)
		differences := diff(anime, malAnime)

		// Add if there were any differences
		if len(differences) > 0 {
			comparisons = append(comparisons, &utils.MALComparison{
				Anime:       anime,
				MALAnime:    malAnime,
				Differences: differences,
			})

			if len(comparisons) >= maxCompareMALEntries {
				break
			}
		}
	}

	return comparisons
}

// diff returns all the differences between an anime and its MAL counterpart.
func diff(anime *arn.Anime, malAnime *mal.Anime) []animediff.Difference {
	// Prealloc linter would complain, but this is best left as nil by default.
	// nolint:prealloc
	var differences []animediff.Difference

	// We'll use the following diffs
	diffFunctions := []diffFunction{
		diffTitles,
		diffDates,
		diffEpisodes,
		diffStatus,
		diffSynopsis,
		diffGenres,
	}

	for _, diffFunction := range diffFunctions {
		diffs := diffFunction(anime, malAnime)
		differences = append(differences, diffs...)
	}

	return differences
}
