package character

import (
	"net/http"
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

const maxDescriptionLength = 170

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

	// Quotes
	quotes := arn.FilterQuotes(func(quote *arn.Quote) bool {
		return !quote.IsDraft && len(quote.Description) > 0 && quote.CharacterID == character.ID
	})

	arn.SortQuotesPopularFirst(quotes)

	// Set OpenGraph attributes
	description := character.Description

	if len(description) > maxDescriptionLength {
		description = description[:maxDescriptionLength-3] + "..."
	}

	ctx.Data = &arn.OpenGraph{
		Tags: map[string]string{
			"og:title":       character.Name,
			"og:image":       "https:" + character.Image,
			"og:url":         "https://" + ctx.App.Config.Domain + character.Link(),
			"og:site_name":   "notify.moe",
			"og:description": description,

			// The OpenGraph type "profile" is meant for real-life persons but I think it's okay in this context.
			// An alternative would be to use "article" which is mostly used for blog posts and news.
			"og:type": "profile",
		},
		Meta: map[string]string{
			"description": description,
			"keywords":    character.Name + ",anime,character",
		},
	}

	return ctx.HTML(components.CharacterDetails(character, characterAnime, quotes, user))
}
