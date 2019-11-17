package genres

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
)

// Get ...
func Get(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)
	genres := []string{}
	genreToAnime := map[string]*arn.Anime{}

	for _, genre := range arn.Genres {
		if genre == "Hentai" {
			continue
		}

		genres = append(genres, genre)
	}

	allAnime := arn.AllAnime()
	arn.SortAnimeByQuality(allAnime)

	added := 0

	for _, anime := range allAnime {
		for _, genre := range anime.Genres {
			// Skip genre that we don't care about
			if !arn.Contains(genres, genre) {
				continue
			}

			_, exists := genreToAnime[genre]

			if !exists {
				genreToAnime[genre] = anime
				added++
			}
		}

		if added >= len(genres) {
			break
		}
	}

	return ctx.HTML(components.Genres(genres, genreToAnime, user))
}
