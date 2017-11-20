package character

import (
	"net/http"
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Get character.
func Get(ctx *aero.Context) string {
	user := utils.GetUser(ctx)
	id := ctx.Get("id")
	character, err := arn.GetCharacter(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Character not found", err)
	}

	characterAnime := character.Anime()

	sort.Slice(characterAnime, func(i, j int) bool {
		if characterAnime[i].StartDate == "" {
			return false
		}

		if characterAnime[j].StartDate == "" {
			return true
		}

		return characterAnime[i].StartDate < characterAnime[j].StartDate
	})

	return ctx.HTML(components.CharacterDetails(character, characterAnime, user))
}
