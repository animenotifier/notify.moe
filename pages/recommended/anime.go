package recommended

import (
	"math"
	"net/http"
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

const (
	maxRecommendations                = 10
	bestGenreCount                    = 3
	genreBonusCompletedAnimeThreshold = 40
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
	completed := animeList.FilterStatus(arn.AnimeListStatusCompleted)

	// Genre affinity
	genreItems := animeList.Genres()
	genreAffinity := map[string]float64{}
	bestGenres := []string{}

	for genre, animeListItems := range genreItems {
		affinity := 0.0

		for _, item := range animeListItems {
			if item.Status != arn.AnimeListStatusCompleted {
				continue
			}

			if item.Rating.Overall != 0 {
				affinity += item.Rating.Overall
			} else {
				affinity += 5.0
			}
		}

		genreAffinity[genre] = affinity
		bestGenres = append(bestGenres, genre)
	}

	sort.Slice(bestGenres, func(i, j int) bool {
		return genreAffinity[bestGenres[i]] > genreAffinity[bestGenres[j]]
	})

	if len(bestGenres) > bestGenreCount {
		bestGenres = bestGenres[:bestGenreCount]
	}

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

		if anime.Type == "tv" {
			tv = append(tv, anime)
		} else if anime.Type == "movie" {
			movies = append(movies, anime)
		}

		// Skip anime from my list (except planned anime)
		existing := animeList.Find(anime.ID)

		if existing != nil && existing.Status != arn.AnimeListStatusPlanned {
			continue
		}

		animeAffinity := anime.Score()

		// Planned anime go higher
		if existing != nil && existing.Status == arn.AnimeListStatusPlanned {
			animeAffinity += 10.0
		}

		// Anime whose high-ranked prequel you did not see are lower ranked
		prequels := anime.Prequels()

		for _, prequel := range prequels {
			item := completed.Find(prequel.ID)

			// Filter out unimportant prequels
			if prequel.Score() < anime.Score()/2 {
				continue
			}

			if item == nil {
				animeAffinity -= 20.0
			}
		}

		// Give favorite genre bonus if we have enough completed anime
		if len(completed.Items) >= genreBonusCompletedAnimeThreshold {
			bestGenreCount := 0

			for _, genre := range anime.Genres {
				if arn.Contains(bestGenres, genre) {
					bestGenreCount++
				}
			}

			// Use square root to dampen the bonus of additional best genres
			animeAffinity += math.Sqrt(float64(bestGenreCount)) * 7.0
		}

		affinity[anime.ID] = animeAffinity
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
