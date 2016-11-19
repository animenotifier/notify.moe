package genre

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

// Get ...
func Get(ctx *aero.Context) string {
	genreName := ctx.Get("name")
	genreInfo := new(arn.Genre)

	err := arn.GetObject("Genres", genreName, genreInfo)

	if err != nil {
		return err.Error()
	}

	return ctx.HTML(components.Genre(genreInfo.Genre, genreInfo.AnimeList))
}
