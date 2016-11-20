package genres

import (
	"fmt"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

// Get ...
func Get(ctx *aero.Context) string {
	var genres []*arn.Genre

	for _, genreName := range arn.Genres {
		// Skip this genre because it doesn't get processed in the background jobs
		if genreName == "Hentai" {
			continue
		}

		genre, err := arn.GetGenre(arn.GetGenreIDByName(genreName))

		if err == nil {
			genre.Name = genreName
			genres = append(genres, genre)
		} else {
			fmt.Println(err)
		}
	}

	return ctx.HTML(components.Genres(genres))
}
