package filteranime

import (
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

	arn.SortAnimeByQuality(missing)

	count := len(missing)

	if count > maxShoboiEntries {
		missing = missing[:maxShoboiEntries]
	}

	return ctx.HTML(components.AnimeEditorListFull(
		"Anime without Shoboi mappings",
		missing,
		count,
		ctx.URI(),
		func(anime *arn.Anime) string {
			return "http://cal.syoboi.jp/find?type=quick&sd=1&kw=" + anime.Title.Japanese
		},
	))
}
