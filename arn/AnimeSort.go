package arn

import (
	"fmt"
	"sort"
	"time"
)

const (
	currentlyAiringBonus      = 5.0
	longSummaryBonus          = 0.1
	popularityThreshold       = 5
	popularityPenalty         = 8.0
	watchingPopularityWeight  = 0.07
	completedPopularityWeight = watchingPopularityWeight
	plannedPopularityWeight   = watchingPopularityWeight * (2.0 / 3.0)
	droppedPopularityWeight   = -plannedPopularityWeight
	visualsWeight             = 0.0075
	storyWeight               = 0.0075
	soundtrackWeight          = 0.0075
	movieBonus                = 0.28
	agePenalty                = 11.0
	ageThreshold              = 6 * 30 * 24 * time.Hour
)

// SortAnimeByPopularity sorts the given slice of anime by popularity.
func SortAnimeByPopularity(animes []*Anime) {
	sort.Slice(animes, func(i, j int) bool {
		aPopularity := animes[i].Popularity.Total()
		bPopularity := animes[j].Popularity.Total()

		if aPopularity == bPopularity {
			return animes[i].Title.Canonical < animes[j].Title.Canonical
		}

		return aPopularity > bPopularity
	})
}

// SortAnimeByQuality sorts the given slice of anime by quality.
func SortAnimeByQuality(animes []*Anime) {
	SortAnimeByQualityDetailed(animes, "")
}

// SortAnimeByQualityDetailed sorts the given slice of anime by quality.
func SortAnimeByQualityDetailed(animes []*Anime, filterStatus string) {
	sort.Slice(animes, func(i, j int) bool {
		a := animes[i]
		b := animes[j]

		scoreA := a.Score()
		scoreB := b.Score()

		// If we show currently running shows, rank shows that started a long time ago a bit lower
		if filterStatus == "current" {
			if a.StartDate != "" && time.Since(a.StartDateTime()) > ageThreshold {
				scoreA -= agePenalty
			}

			if b.StartDate != "" && time.Since(b.StartDateTime()) > ageThreshold {
				scoreB -= agePenalty
			}
		}

		if scoreA == scoreB {
			return a.Title.Canonical < b.Title.Canonical
		}

		return scoreA > scoreB
	})
}

// Score returns the score used for the anime ranking.
func (anime *Anime) Score() float64 {
	score := anime.Rating.Overall
	score += anime.Rating.Story * storyWeight
	score += anime.Rating.Visuals * visualsWeight
	score += anime.Rating.Soundtrack * soundtrackWeight

	score += float64(anime.Popularity.Watching) * watchingPopularityWeight
	score += float64(anime.Popularity.Planned) * plannedPopularityWeight
	score += float64(anime.Popularity.Completed) * completedPopularityWeight
	score += float64(anime.Popularity.Dropped) * droppedPopularityWeight

	if anime.Status == "current" {
		score += currentlyAiringBonus
	}

	if anime.Type == "movie" {
		score += movieBonus
	}

	if anime.Popularity.Total() < popularityThreshold {
		score -= popularityPenalty
	}

	if len(anime.Summary) >= 140 {
		score += longSummaryBonus
	}

	return score
}

// ScoreHumanReadable returns the score used for the anime ranking in human readable format.
func (anime *Anime) ScoreHumanReadable() string {
	return fmt.Sprintf("%.1f", anime.Score())
}
