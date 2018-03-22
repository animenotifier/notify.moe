package filteranime

import (
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

	arn.SortAnimeByQuality(missing)

	count := len(missing)

	if count > maxAniListEntries {
		missing = missing[:maxAniListEntries]
	}

	return ctx.HTML(components.AnimeEditorListFull(
		"Anime without Anilist mappings",
		missing,
		count,
		ctx.URI(),
		func(anime *arn.Anime) string {
			return "https://anilist.co/search?type=anime&q=" + anime.Title.Canonical
		},
	))
}
