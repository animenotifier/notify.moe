package recommended

import (
	"math"

	"github.com/animenotifier/notify.moe/arn"
)

func getAnimeAffinity(anime *arn.Anime, animeListItem *arn.AnimeListItem, completed *arn.AnimeList, bestGenres []string) float64 {
	animeAffinity := anime.Score()

	// Planned anime go higher
	if animeListItem != nil && animeListItem.Status == arn.AnimeListStatusPlanned {
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

	return animeAffinity
}
