package main

import (
	"fmt"

	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
)

var ratings = map[string][]arn.AnimeListItemRating{}
var finalRating = map[string]*arn.AnimeRating{}
var popularity = map[string]*arn.AnimePopularity{}

// Note this is using the airing-anime as a template with modfications
// made to it.
func main() {
	color.Yellow("Updating anime ratings")
	color.Cyan(arn.Node.Address().String())

	defer color.Green("Finished.")
	defer arn.Node.Close()

	fmt.Println("Processing anime lists")

	for animeList := range arn.StreamAnimeLists() {
		extractRatings(animeList)
		extractPopularity(animeList)
	}

	// Calculate rating
	for animeID := range finalRating {
		overall := []float64{}
		story := []float64{}
		visuals := []float64{}
		soundtrack := []float64{}

		for _, rating := range ratings[animeID] {
			if rating.Overall != 0 {
				overall = append(overall, rating.Overall)
			}

			if rating.Story != 0 {
				story = append(story, rating.Story)
			}

			if rating.Visuals != 0 {
				visuals = append(visuals, rating.Visuals)
			}

			if rating.Soundtrack != 0 {
				soundtrack = append(soundtrack, rating.Soundtrack)
			}
		}

		// Save number of people who rated on this
		finalRating[animeID].Count.Overall = len(overall)
		finalRating[animeID].Count.Story = len(story)
		finalRating[animeID].Count.Visuals = len(visuals)
		finalRating[animeID].Count.Soundtrack = len(soundtrack)

		// Dampen the rating if number of users is too low
		if len(overall) < arn.RatingCountThreshold {
			overall = append(overall, arn.AverageRating)
		}

		if len(story) < arn.RatingCountThreshold {
			story = append(story, arn.AverageRating)
		}

		if len(visuals) < arn.RatingCountThreshold {
			visuals = append(visuals, arn.AverageRating)
		}

		if len(soundtrack) < arn.RatingCountThreshold {
			soundtrack = append(soundtrack, arn.AverageRating)
		}

		// Average rating
		finalRating[animeID].Overall = average(overall)
		finalRating[animeID].Story = average(story)
		finalRating[animeID].Visuals = average(visuals)
		finalRating[animeID].Soundtrack = average(soundtrack)
	}

	// Save rating
	fmt.Println("Saving rating")

	for animeID := range finalRating {
		anime, err := arn.GetAnime(animeID)

		if err != nil {
			panic(err)
		}

		anime.Rating = *finalRating[animeID]
		anime.Save()
	}

	// Save popularity
	fmt.Println("Saving popularity")

	for animeID := range popularity {
		anime, err := arn.GetAnime(animeID)

		if err != nil {
			panic(err)
		}

		anime.Popularity = *popularity[animeID]
		anime.Save()
	}
}

func average(floatSlice []float64) float64 {
	if len(floatSlice) == 0 {
		return 0
	}

	var sum float64

	for _, value := range floatSlice {
		sum += value
	}

	return sum / float64(len(floatSlice))
}

func extractRatings(animeList *arn.AnimeList) {
	for _, item := range animeList.Items {
		if item.Rating.IsNotRated() {
			continue
		}

		_, found := ratings[item.AnimeID]

		if !found {
			ratings[item.AnimeID] = []arn.AnimeListItemRating{}
			finalRating[item.AnimeID] = &arn.AnimeRating{}
		}

		ratings[item.AnimeID] = append(ratings[item.AnimeID], item.Rating)
	}
}

func extractPopularity(animeList *arn.AnimeList) {
	for _, item := range animeList.Items {
		_, found := popularity[item.AnimeID]

		if !found {
			popularity[item.AnimeID] = &arn.AnimePopularity{}
		}

		counter := popularity[item.AnimeID]

		switch item.Status {
		case arn.AnimeListStatusWatching:
			counter.Watching++
		case arn.AnimeListStatusCompleted:
			counter.Completed++
		case arn.AnimeListStatusPlanned:
			counter.Planned++
		case arn.AnimeListStatusHold:
			counter.Hold++
		case arn.AnimeListStatusDropped:
			counter.Dropped++
		}
	}
}
