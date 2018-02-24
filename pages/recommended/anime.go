package recommended

import (
	"net/http"
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

const (
	maxRecommendations = 20
	worstGenreCount    = 5
)

// Anime shows a list of recommended anime.
func Anime(ctx *aero.Context) string {
	user := utils.GetUser(ctx)
	nick := ctx.Get("nick")
	viewUser, err := arn.GetUserByNick(nick)

	if err != nil {
		return ctx.Error(http.StatusUnauthorized, "Not logged in", err)
	}

	animeList := viewUser.AnimeList()
	genreItems := animeList.Genres()
	genreAffinity := map[string]float64{}
	worstGenres := []string{}

	for genre, animeListItems := range genreItems {
		affinity := 0.0

		for _, item := range animeListItems {
			// if item.Status == arn.AnimeListStatusDropped {
			// 	affinity -= 5.0
			// 	continue
			// }

			if item.Rating.Overall != 0 {
				affinity += item.Rating.Overall
			} else {
				affinity += 5.0
			}
		}

		genreAffinity[genre] = affinity
		worstGenres = append(worstGenres, genre)
	}

	sort.Slice(worstGenres, func(i, j int) bool {
		return genreAffinity[worstGenres[i]] < genreAffinity[worstGenres[j]]
	})

	if len(worstGenres) > worstGenreCount {
		worstGenres = worstGenres[:worstGenreCount]
	}

	// Get all anime
	recommendations := arn.AllAnime()

	// Affinity maps an anime ID to a number that indicates how likely a user is going to enjoy that anime.
	affinity := map[string]float64{}

	// Calculate affinity for each anime
	for _, anime := range recommendations {
		// Skip anime that are upcoming
		if anime.Status == "upcoming" {
			continue
		}

		// Skip anime from my list (except planned anime)
		existing := animeList.Find(anime.ID)

		if existing != nil && existing.Status != arn.AnimeListStatusPlanned {
			continue
		}

		// Skip anime that don't have one of the top genres for that user
		worstGenreFound := false

		for _, genre := range anime.Genres {
			if arn.Contains(worstGenres, genre) {
				worstGenreFound = true
				break
			}
		}

		if worstGenreFound {
			continue
		}

		animeAffinity := 0.0

		// Planned anime go higher
		if existing != nil && existing.Status == arn.AnimeListStatusPlanned {
			animeAffinity += 75.0
		}

		animeAffinity += float64(anime.Popularity.Total())
		affinity[anime.ID] = animeAffinity
	}

	// Sort
	sort.Slice(recommendations, func(i, j int) bool {
		affinityA := affinity[recommendations[i].ID]
		affinityB := affinity[recommendations[j].ID]

		if affinityA == affinityB {
			return recommendations[i].Title.Canonical < recommendations[j].Title.Canonical
		}

		return affinityA > affinityB
	})

	// Take the top 10
	if len(recommendations) > maxRecommendations {
		recommendations = recommendations[:maxRecommendations]
	}

	return ctx.HTML(components.RecommendedAnime(recommendations, worstGenres, viewUser, user))
}
