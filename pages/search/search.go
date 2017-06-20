package search

import (
	"strings"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

const maxUsers = 9 * 7
const maxAnime = 9 * 7

// Get search page.
func Get(ctx *aero.Context) string {
	term := strings.ToLower(ctx.Get("term"))

	var users []*arn.User
	var animeResults []*arn.Anime

	// Search everything in parallel
	aero.Parallel(func() {
		// Search users
		var user *arn.User

		userSearchIndex, err := arn.GetSearchIndex("User")

		if err != nil {
			return
		}

		for name, id := range userSearchIndex.TextToID {
			if strings.Index(name, term) != -1 {
				user, err = arn.GetUser(id)

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
		// Search anime
		var anime *arn.Anime

		animeSearchIndex, err := arn.GetSearchIndex("Anime")

		if err != nil {
			return
		}

		for title, id := range animeSearchIndex.TextToID {
			if strings.Index(title, term) != -1 {
				anime, err = arn.GetAnime(id)

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
