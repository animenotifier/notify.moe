package editor

import (
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

const maxShoboiEntries = 70

// Shoboi ...
func Shoboi(ctx *aero.Context) string {
	missing := arn.FilterAnime(func(anime *arn.Anime) bool {
		return anime.GetMapping("shoboi/anime") == ""
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

	if len(missing) > maxShoboiEntries {
		missing = missing[:maxShoboiEntries]
	}

	return ctx.HTML(components.ShoboiMissingMapping(missing))
}
