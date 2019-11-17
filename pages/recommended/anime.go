package recommended

import (
	"net/http"
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
)

const (
	maxRecommendations                = 10
	bestGenreCount                    = 3
	genreBonusCompletedAnimeThreshold = 40
)

// Anime shows a list of recommended anime.
func Anime(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)
	nick := ctx.Get("nick")
	viewUser, err := arn.GetUserByNick(nick)

	if err != nil {
		return ctx.Error(http.StatusUnauthorized, "Not logged in", err)
	}

	animeList := viewUser.AnimeList()
	completed := animeList.FilterStatus(arn.AnimeListStatusCompleted)

	// Genre affinity
	bestGenres := animeList.TopGenres(bestGenreCount)

	// Get all anime
	var tv []*arn.Anime
	var movies []*arn.Anime
	allAnime := arn.AllAnime()

	// Affinity maps an anime ID to a number that indicates how likely a user is going to enjoy that anime.
	affinity := map[string]float64{}

	// Calculate affinity for each anime
	for _, anime := range allAnime {
		// Skip anime that are upcoming or tba
		if anime.Status == "upcoming" || anime.Status == "tba" {
			continue
		}

		switch anime.Type {
		case "tv":
			tv = append(tv, anime)
		case "movie":
			movies = append(movies, anime)
		default:
			continue
		}

		// Skip anime from my list (except planned anime)
		existing := animeList.Find(anime.ID)

		if existing != nil && existing.Status != arn.AnimeListStatusPlanned {
			continue
		}

		affinity[anime.ID] = getAnimeAffinity(anime, existing, completed, bestGenres)
	}

	// Sort
	sort.Slice(tv, func(i, j int) bool {
		affinityA := affinity[tv[i].ID]
		affinityB := affinity[tv[j].ID]

		if affinityA == affinityB {
			return tv[i].Title.Canonical < tv[j].Title.Canonical
		}

		return affinityA > affinityB
	})

	sort.Slice(movies, func(i, j int) bool {
		affinityA := affinity[movies[i].ID]
		affinityB := affinity[movies[j].ID]

		if affinityA == affinityB {
			return movies[i].Title.Canonical < movies[j].Title.Canonical
		}

		return affinityA > affinityB
	})

	// Take the top 10
	if len(tv) > maxRecommendations {
		tv = tv[:maxRecommendations]
	}

	if len(movies) > maxRecommendations {
		movies = movies[:maxRecommendations]
	}

	return ctx.HTML(components.RecommendedAnime(tv, movies, viewUser, user))
}
