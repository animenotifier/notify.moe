package arn

import (
	"sort"
)

// GetCompanyToAnimeMap returns a map that contains company IDs as keys and their anime as values.
func GetCompanyToAnimeMap() map[string][]*Anime {
	companyToAnimes := map[string][]*Anime{}

	allAnime := AllAnime()
	SortAnimeByQuality(allAnime)

	for _, anime := range allAnime {
		for _, studioID := range anime.StudioIDs {
			companyToAnimes[studioID] = append(companyToAnimes[studioID], anime)
		}
	}

	return companyToAnimes
}

// SortCompaniesPopularFirst ...
func SortCompaniesPopularFirst(companies []*Company) {
	// Generate company ID to popularity map
	popularity := map[string]int{}

	for anime := range StreamAnime() {
		for _, studio := range anime.Studios() {
			popularity[studio.ID] += anime.Popularity.Watching + anime.Popularity.Completed
		}
	}

	// Sort by using the popularity map
	sort.Slice(companies, func(i, j int) bool {
		a := companies[i]
		b := companies[j]

		aPopularity := popularity[a.ID]
		bPopularity := popularity[b.ID]

		if aPopularity == bPopularity {
			return a.Name.English < b.Name.English
		}

		return aPopularity > bPopularity
	})
}
