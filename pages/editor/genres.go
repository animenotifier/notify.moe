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
	missing := arn.FilterAnime(func(anime *arn.Anime) bool {
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

	if len(missing) > maxGenreEntries {
		missing = missing[:maxGenreEntries]
	}

	return ctx.HTML(components.AnimeWithoutGenres(missing, nil))
}
