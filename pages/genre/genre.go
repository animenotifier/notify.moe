package genre

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Get ...
func Get(ctx *aero.Context) string {
	user := utils.GetUser(ctx)
	genreName := ctx.Get("name")
	genre, err := arn.GetGenre(genreName)

	if err != nil {
		return ctx.Error(404, "Genre not found", err)
	}

	return ctx.HTML(components.Genre(genre, user))
}
