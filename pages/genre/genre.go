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
	mean := 0.0

	var realGenreName string

	for _, realGenre := range arn.Genres {
		if strings.ToLower(realGenre) == genreName {
			realGenreName = realGenre
			break
		}
	}

	completedItems := user.AnimeList().FilterStatus(arn.AnimeListStatusCompleted).Genres()[realGenreName]
	completed := len(completedItems)

	genreItems := user.AnimeList().Genres()[realGenreName]
	for _, item := range genreItems {
		if item.Rating.IsNotRated() {
			continue
		}
		mean += item.Rating.Overall
	}

	mean = mean / float64(len(genreItems))

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

// containsLowerCase tells you whether the given element exists when all elements are lowercased.
func containsLowerCase(array []string, search string) bool {
	for _, item := range array {
		if strings.ToLower(item) == search {
			return true
		}
	}

	return false
}
