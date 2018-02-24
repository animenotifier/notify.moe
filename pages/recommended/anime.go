package recommended

import (
	"net/http"
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

const maxRecommendations = 20

// Anime shows a list of recommended anime.
func Anime(ctx *aero.Context) string {
	nick := ctx.Get("nick")
	user, err := arn.GetUserByNick(nick)

	if err != nil {
		return ctx.Error(http.StatusUnauthorized, "Not logged in", err)
	}

	animeList := user.AnimeList()
	genreItems := animeList.Genres()
	genreAffinity := map[string]float64{}

	for genre, animeListItems := range genreItems {
		affinity := 0.0

		for _, item := range animeListItems {
			if item.Status == arn.AnimeListStatusDropped {
				affinity -= 5.0
				continue
			}

			if item.Rating.Overall != 0 {
				affinity += item.Rating.Overall
			} else {
				affinity += 5.0
			}
		}

		genreAffinity[genre] = affinity
	}

	// Get all anime
	recommendations := arn.AllAnime()

	// Affinity maps an anime ID to a number that indicates how likely a user is going to enjoy that anime.
	affinity := map[string]float64{}

	// Calculate affinity for each anime
	for _, anime := range recommendations {
		// Skip anime from my list (except planned anime)
		existing := animeList.Find(anime.ID)

		if existing != nil && existing.Status != arn.AnimeListStatusPlanned {
			continue
		}

		affinity[anime.ID] = float64(anime.Popularity.Total())

		// animeGenresAffinity := 0.0

		// if len(anime.Genres) > 0 {
		// 	for _, genre := range anime.Genres {
		// 		if genreAffinity[genre] > animeGenresAffinity {
		// 			animeGenresAffinity = genreAffinity[genre]
		// 		}
		// 	}

		// 	animeGenresAffinity = animeGenresAffinity / float64(len(anime.Genres))
		// }

		// affinity[anime.ID] = animeGenresAffinity
	}

	// Sort
	sort.Slice(recommendations, func(i, j int) bool {
		return affinity[recommendations[i].ID] > affinity[recommendations[j].ID]
	})

	// Take the top 10
	if len(recommendations) > maxRecommendations {
		recommendations = recommendations[:maxRecommendations]
	}

	return ctx.HTML(components.RecommendedAnime(recommendations, user))
}
