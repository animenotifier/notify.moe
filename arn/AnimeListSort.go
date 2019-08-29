package arn

import "sort"

// Sort sorts the anime list by the given algorithm.
func (list *AnimeList) Sort(algorithm string) {
	list.Lock()
	defer list.Unlock()

	switch algorithm {
	case SortByTitle:
		sort.Slice(list.Items, func(i, j int) bool {
			a := list.Items[i]
			b := list.Items[j]

			return a.Anime().Title.Canonical < b.Anime().Title.Canonical
		})

	case SortByRating:
		sort.Slice(list.Items, func(i, j int) bool {
			a := list.Items[i]
			b := list.Items[j]

			if a.Rating.Overall == b.Rating.Overall {
				return a.Anime().Title.Canonical < b.Anime().Title.Canonical
			}

			return a.Rating.Overall > b.Rating.Overall
		})

	case SortByAiringDate:
		sort.Slice(list.Items, func(i, j int) bool {
			a := list.Items[i]
			b := list.Items[j]

			epsA := a.Anime().UpcomingEpisode()
			epsB := b.Anime().UpcomingEpisode()

			if epsA == nil && epsB == nil {
				if a.Rating.Overall == b.Rating.Overall {
					return a.Anime().Title.Canonical < b.Anime().Title.Canonical
				}

				return a.Rating.Overall > b.Rating.Overall
			}

			if epsA == nil {
				return false
			}

			if epsB == nil {
				return true
			}

			return epsA.Episode.AiringDate.Start < epsB.Episode.AiringDate.Start
		})
	}
}
