package editor

import (
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

const maxAniListEntries = 70

// AniList ...
func AniList(ctx *aero.Context) string {
	year, _ := ctx.GetInt("year")
	animeType := ctx.Get("type")

	missing := arn.FilterAnime(func(anime *arn.Anime) bool {
		if year != 0 && year != anime.StartDateTime().Year() {
			return false
		}

		if animeType != "" && anime.Type != animeType {
			return false
		}

		return anime.GetMapping("anilist/anime") == ""
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

	if count > maxAniListEntries {
		missing = missing[:maxAniListEntries]
	}

	return ctx.HTML(components.AnimeEditorListFull(
		"Anime without Anilist links",
		missing,
		count,
		ctx.URI(),
		func(anime *arn.Anime) string {
			return "https://anilist.co/search?type=anime&q=" + anime.Title.Canonical
		},
	))
}
