package filteranime

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

const maxAnimeEntries = 70

// editorList renders the anime list with the given title and filter.
func editorList(ctx *aero.Context, title string, filter func(*arn.Anime) bool, searchLink func(*arn.Anime) string) string {
	animes, count := filterAnime(ctx, filter)

	return ctx.HTML(components.AnimeEditorListFull(
		title,
		animes,
		count,
		ctx.URI(),
		searchLink,
	))
}

// filterAnime filters anime by the given filter function and
// additionally applies year and types filters if specified.
func filterAnime(ctx *aero.Context, filter func(*arn.Anime) bool) ([]*arn.Anime, int) {
	year, _ := ctx.GetInt("year")
	animeType := ctx.Get("type")

	// Filter
	animes := arn.FilterAnime(func(anime *arn.Anime) bool {
		if year != 0 && year != anime.StartDateTime().Year() {
			return false
		}

		if animeType != "" && anime.Type != animeType {
			return false
		}

		return filter(anime)
	})

	// Sort
	arn.SortAnimeByQuality(animes)

	// Limit
	count := len(animes)

	if count > maxAnimeEntries {
		animes = animes[:maxAnimeEntries]
	}

	return animes, count
}
