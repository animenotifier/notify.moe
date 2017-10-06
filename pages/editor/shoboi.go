package editor

import (
	"net/http"
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

const maxShoboiEntries = 70

// Shoboi ...
func Shoboi(ctx *aero.Context) string {
	missing, err := arn.FilterAnime(func(anime *arn.Anime) bool {
		return anime.GetMapping("shoboi/anime") == ""
	})

	if err != nil {
		ctx.Error(http.StatusInternalServerError, "Couldn't filter anime", err)
	}

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
