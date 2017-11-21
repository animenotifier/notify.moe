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

	// Set OpenGraph attributes
	ctx.Data = &arn.OpenGraph{
		Tags: map[string]string{
			"og:title":       character.Name,
			"og:image":       character.Image,
			"og:url":         "https://" + ctx.App.Config.Domain + character.Link(),
			"og:site_name":   "notify.moe",
			"og:description": character.Description,
		},
		Meta: map[string]string{
			"description": character.Description,
			"keywords":    character.Name + ",anime,character",
		},
	}

	return ctx.HTML(components.CharacterDetails(character, characterAnime, user))
}
