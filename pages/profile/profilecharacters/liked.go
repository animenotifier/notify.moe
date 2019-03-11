package profilecharacters

import (
	"net/http"
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Liked shows all liked characters of a particular user.
func Liked(ctx *aero.Context) string {
	nick := ctx.Get("nick")
	viewUser, err := arn.GetUserByNick(nick)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "User not found", err)
	}

	characters := []*arn.Character{}

	for character := range arn.StreamCharacters() {
		if arn.Contains(character.Likes, viewUser.ID) {
			characters = append(characters, character)
		}
	}

	sort.Slice(characters, func(i, j int) bool {
		return characters[i].Name.Canonical < characters[j].Name.Canonical

		// aLikes := len(characters[i].Likes)
		// bLikes := len(characters[j].Likes)

		// if aLikes == bLikes {
		// 	return characters[i].Name.Canonical < characters[j].Name.Canonical
		// }

		// return aLikes > bLikes
	})

	return ctx.HTML(components.ProfileCharacters(characters, viewUser, utils.GetUser(ctx), ctx.URI()))
}
