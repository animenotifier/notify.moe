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
	mainQuote := character.MainQuote()
	quotes := character.Quotes()

	arn.SortQuotesPopularFirst(quotes)

	// Set OpenGraph attributes
	description := character.Description

	if len(description) > maxDescriptionLength {
		description = description[:maxDescriptionLength-3] + "..."
	}

	ctx.Data = &arn.OpenGraph{
		Tags: map[string]string{
			"og:title":       character.Name.Canonical,
			"og:image":       "https:" + character.ImageLink("large"),
			"og:url":         "https://" + ctx.App.Config.Domain + character.Link(),
			"og:site_name":   "notify.moe",
			"og:description": description,

			// The OpenGraph type "profile" is meant for real-life persons but I think it's okay in this context.
			// An alternative would be to use "article" which is mostly used for blog posts and news.
			"og:type": "profile",
		},
		Meta: map[string]string{
			"description": description,
			"keywords":    character.Name.Canonical + ",anime,character",
		},
	}

	// Friends
	var friends []*arn.User

	if user != nil {
		friendIDs := utils.Intersection(character.Likes, user.Follows().Items)
		friendObjects := arn.DB.GetMany("User", friendIDs)

		for _, obj := range friendObjects {
			if obj == nil {
				continue
			}

			friends = append(friends, obj.(*arn.User))
		}
	}

	return ctx.HTML(components.CharacterDetails(character, characterAnime, quotes, friends, mainQuote, user))
}
