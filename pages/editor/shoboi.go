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
	year, _ := ctx.GetInt("year")
	animeType := ctx.Get("type")

	missing := arn.FilterAnime(func(anime *arn.Anime) bool {
		if year != 0 && year != anime.StartDateTime().Year() {
			return false
		}

		if animeType != "" && anime.Type != animeType {
			return false
		}

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

	count := len(missing)

	if count > maxShoboiEntries {
		missing = missing[:maxShoboiEntries]
	}

	return ctx.HTML(components.AnimeEditorListFull(
		"Anime without Shoboi links",
		missing,
		count,
		"/editor/anime/missing/shoboi",
		func(anime *arn.Anime) string {
			return "http://cal.syoboi.jp/find?type=quick&sd=1&kw=" + anime.Title.Japanese
		},
	))
}
