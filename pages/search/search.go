package search

import (
	"strings"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

const maxUsers = 18
const maxAnime = 18

type AnimeID = string
type UserID = string

var animeSearchIndex = make(map[string]AnimeID)
var userSearchIndex = make(map[string]UserID)

func init() {
	updateSearchIndex()
}

func updateSearchIndex() {
	updateAnimeIndex()
	updateUserIndex()
}

func updateAnimeIndex() {
	// Anime
	animeStream, err := arn.AllAnime()

	if err != nil {
		panic(err)
	}
	
	for anime := range animeStream {
		animeSearchIndex[strings.ToLower(anime.Title.Canonical)] = anime.ID
	}
}

func updateUserIndex() {
	// Users
	userStream, err := arn.AllUsers()

	if err != nil {
		panic(err)
	}

	for user := range userStream {
		userSearchIndex[strings.ToLower(user.Nick)] = user.ID
	}
}

// Get search page.
func Get(ctx *aero.Context) string {
	term := strings.ToLower(ctx.Get("term"))

	var users []*arn.User
	var animeResults []*arn.Anime

	aero.Parallel(func() {
		for name, id := range userSearchIndex {
			if strings.Index(name, term) != -1 {
				user, err := arn.GetUser(id)

				if err != nil {
					continue
				}

				users = append(users, user)

				if len(users) >= maxUsers {
					break
				}
			}
		}
	}, func() {
		for title, id := range animeSearchIndex {
			if strings.Index(title, term) != -1 {
				anime, err := arn.GetAnime(id)

				if err != nil {
					continue
				}

				animeResults = append(animeResults, anime)

				if len(animeResults) >= maxAnime {
					break
				}
			}
		}
	})

	return ctx.HTML(components.Search(users, animeResults))
}
