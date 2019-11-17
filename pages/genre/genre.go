package genre

import (
	"strings"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
)

const animePerPage = 100
const animeRatingCountThreshold = 5

// Get renders the genre page.
func Get(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)
	genreName := ctx.Get("name")
	animes := []*arn.Anime{}

	for anime := range arn.StreamAnime() {
		if containsLowerCase(anime.Genres, genreName) {
			animes = append(animes, anime)
		}
	}

	userScore := averageUserScore(user, animes)
	userCompleted := totalCompleted(user, animes)
	globalScore := averageGlobalScore(animes)

	arn.SortAnimeByQuality(animes)

	if len(animes) > animePerPage {
		animes = animes[:animePerPage]
	}

	return ctx.HTML(components.Genre(genreName, animes, user, userScore, userCompleted, globalScore))
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

// averageUserScore counts the user's average score for a list of anime.
func averageUserScore(user *arn.User, animes []*arn.Anime) float64 {
	if user == nil {
		return 0
	}

	count := 0.0
	scores := 0.0

	animeList := user.AnimeList()

	for _, anime := range animes {
		userAnime := animeList.Find(anime.ID)

		if userAnime != nil && !userAnime.Rating.IsNotRated() {
			scores += userAnime.Rating.Overall
			count++
		}
	}

	if count == 0.0 {
		return 0
	}

	return scores / count
}

// averageGlobalScore returns the average overall score for the given anime.
func averageGlobalScore(animes []*arn.Anime) float64 {
	sum := 0.0
	count := 0

	for _, anime := range animes {
		if anime.Rating.Count.Overall >= animeRatingCountThreshold {
			sum += anime.Rating.Overall
			count++
		}
	}

	return sum / float64(count)
}

// totalCompleted counts the number of anime the user has completed from a given list of anime.
func totalCompleted(user *arn.User, animes []*arn.Anime) int {
	if user == nil {
		return 0
	}

	count := 0

	completedList := user.AnimeList().FilterStatus(arn.AnimeListStatusCompleted)

	for _, anime := range animes {
		if completedList.Contains(anime.ID) {
			count++
		}
	}

	return count
}
