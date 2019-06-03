package filteranime

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
)

// AniList ...
func AniList(ctx aero.Context) error {
	return editorList(
		ctx,
		"Anime without Anilist mappings",
		func(anime *arn.Anime) bool {
			return anime.GetMapping("anilist/anime") == ""
		},
		func(anime *arn.Anime) string {
			return "https://anilist.co/search/anime?sort=SEARCH_MATCH&search=" + anime.Title.Canonical
		},
	)
}
