package admin

import (
	"net/http"
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

// AniList ...
func AniList(ctx *aero.Context) string {
	missing, err := arn.FilterAnime(func(anime *arn.Anime) bool {
		return anime.GetMapping("anilist/anime") == ""
	})

	if err != nil {
		ctx.Error(http.StatusInternalServerError, "Couldn't filter anime", err)
	}

	sort.Slice(missing, func(i, j int) bool {
		return missing[i].StartDate > missing[j].StartDate
	})

	return ctx.HTML(components.AniListMissingMapping(missing))
}
