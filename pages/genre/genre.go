package genre

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

// Get ...
func Get(ctx *aero.Context) string {
	genreName := ctx.Get("name")
	genre, err := arn.GetGenre(genreName)

	if err != nil {
		return err.Error()
	}

	return ctx.HTML(components.Genre(genre))
}
