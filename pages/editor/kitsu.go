package editor

import (
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/kitsu"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
)

// NewKitsuAnime ...
func NewKitsuAnime(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)
	finder := arn.NewAnimeFinder("kitsu/anime")
	deletedIDs, err := arn.GetIDList("deleted kitsu anime")

	if err != nil {
		deletedIDs = arn.IDList{}
	}

	animes := arn.FilterKitsuAnime(func(anime *kitsu.Anime) bool {
		return finder.GetAnime(anime.ID) == nil && !arn.Contains(deletedIDs, anime.ID)
	})

	sort.Slice(animes, func(i, j int) bool {
		a := animes[i]
		b := animes[j]

		return a.ID > b.ID
	})

	return ctx.HTML(components.NewKitsuAnime(animes, ctx.Path(), user))
}
