package editor

import (
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

const maxGenreEntries = 70

// Genres ...
func Genres(ctx *aero.Context) string {
	year, _ := ctx.GetInt("year")
	animeType := ctx.Get("type")

	missing := arn.FilterAnime(func(anime *arn.Anime) bool {
		if year != 0 && year != anime.StartDateTime().Year() {
			return false
		}

		if animeType != "" && anime.Type != animeType {
			return false
		}

		return len(anime.Genres) == 0
	})

	sort.Slice(missing, func(i, j int) bool {
		a := missing[i]
		b := missing[j]

		aPop := a.Popularity.Total()
		bPop := b.Popularity.Total()

		if aPop == bPop {
			return a.Title.Canonical < b.Title.Canonical
		}

		return aPop > bPop
	})

	count := len(missing)

	if count > maxGenreEntries {
		missing = missing[:maxGenreEntries]
	}

	return ctx.HTML(components.AnimeEditorListFull(
		"Anime without genres",
		missing,
		count,
		"/editor/anime/missing/genres",
		nil,
	))
}
