package genres

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

// Get ...
func Get(ctx *aero.Context) string {
	var genres []string

	for _, genre := range arn.Genres {
		if genre == "Hentai" {
			continue
		}

		genres = append(genres, genre)
	}

	return ctx.HTML(components.Genres(genres))
}
