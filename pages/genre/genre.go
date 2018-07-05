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
	animes := []*arn.Anime{}

	for anime := range arn.StreamAnime() {
		if containsLowerCase(anime.Genres, genreName) {
			animes = append(animes, anime)
		}
	}

	userScore := averageScore(user, animes)

	arn.SortAnimeByQuality(animes)

	if len(animes) > animePerPage {
		animes = animes[:animePerPage]
	}

	return ctx.HTML(components.Genre(genreName, animes, user, userScore))
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

// averageScore counts the user's average score for the given animes.
func averageScore(user *arn.User, animes []*arn.Anime) float64 {
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
