package recommended

import (
	"sort"

	"github.com/animenotifier/arn"
)

// getBestGenres returns the most liked genres for the user's anime list.
func getBestGenres(animeList *arn.AnimeList) []string {
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

	return bestGenres
}
