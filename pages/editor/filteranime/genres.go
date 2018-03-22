package filteranime

import (
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

	arn.SortAnimeByQuality(missing)

	count := len(missing)

	if count > maxGenreEntries {
		missing = missing[:maxGenreEntries]
	}

	return ctx.HTML(components.AnimeEditorListFull(
		"Anime without genres",
		missing,
		count,
		ctx.URI(),
		nil,
	))
}
