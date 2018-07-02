package genre

import (
	"strings"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

const animePerPage = 100

// Get ...
func Get(ctx *aero.Context) string {
	user := utils.GetUser(ctx)
	genreName := ctx.Get("name")
	realGenreName := getGenreName(genreName)
	animes := []*arn.Anime{}
	mean := 0.0
	completed := 0

	if user != nil {
		completedItems := user.AnimeList().FilterStatus(arn.AnimeListStatusCompleted).Genres()[realGenreName]
		completed = len(completedItems)
		mean = userAverage(user, realGenreName)
	}

	for anime := range arn.StreamAnime() {
		if containsLowerCase(anime.Genres, genreName) {
			animes = append(animes, anime)
		}
	}

	arn.SortAnimeByQuality(animes)

	if len(animes) > animePerPage {
		animes = animes[:animePerPage]
	}

	return ctx.HTML(components.Genre(genreName, animes, user, mean, completed))
}

// userAverage return the user average score for a genre
func userAverage(user *arn.User, realGenreName string) float64 {
	genreItems := user.AnimeList().Genres()[realGenreName]
	average := 0.0
	for _, item := range genreItems {
		if item.Rating.IsNotRated() {
			continue
		}

		average += item.Rating.Overall
	}

	return average / float64(len(genreItems))
}

// getGenreName return the normally used genre name from it's lowercase counterpart
func getGenreName(genreName string) string {
	for _, realGenreName := range arn.Genres {
		if strings.ToLower(realGenreName) == genreName {
			return realGenreName
		}
	}

	return ""
}

// containsLowerCase tells you whether the given element exists when all elements are lowercased.
func containsLowerCase(array []string, search string) bool {
	for _, item := range array {
		if strings.ToLower(item) == search {
			return true
		}
	}

	return false
}
