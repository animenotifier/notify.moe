package genre

import (
	"strconv"
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

	userScore := averageGenreScore(user, animes)

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

func averageGenreScore(user *arn.User, animes []*arn.Anime) string {
	if user == nil {
		return ""
	}

	counter := 0.0
	scores := 0.0

	for _, anime := range animes {

		if user.AnimeList().Contains(anime.ID) {
			scores = scores + user.AnimeList().Find(anime.ID).Rating.Overall
			counter = counter + 1
		}
	}

	return strconv.FormatFloat(scores/counter, 'f', 6, 64)
}
