package character

import (
	"net/http"
	"sort"

	"github.com/aerogo/aero"
	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/assets"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/server/middleware"
	"github.com/animenotifier/notify.moe/utils"
)

const (
	maxRelevantCharacters = 12
)

// Get character.
func Get(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)
	id := ctx.Get("id")
	character, err := arn.GetCharacter(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Character not found", err)
	}

	// Anime
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

	// Characters from the same anime
	characterAppearances := map[string]int{}

	for _, anime := range characterAnime {
		for _, animeCharacter := range anime.Characters().Items {
			if animeCharacter.CharacterID == character.ID {
				continue
			}

			characterAppearances[animeCharacter.CharacterID]++
		}
	}

	relevantCharacters := []*arn.Character{}

	for characterID := range characterAppearances {
		relevantCharacter, err := arn.GetCharacter(characterID)

		if err != nil {
			color.Red(err.Error())
			continue
		}

		if !relevantCharacter.HasImage() {
			continue
		}

		relevantCharacters = append(relevantCharacters, relevantCharacter)
	}

	sort.Slice(relevantCharacters, func(i, j int) bool {
		aRelevance := characterAppearances[relevantCharacters[i].ID]
		bRelevance := characterAppearances[relevantCharacters[j].ID]

		if aRelevance == bRelevance {
			aLikes := len(relevantCharacters[i].Likes)
			bLikes := len(relevantCharacters[j].Likes)

			if aLikes == bLikes {
				return relevantCharacters[i].Name.Canonical < relevantCharacters[j].Name.Canonical
			}

			return aLikes > bLikes
		}

		return aRelevance > bRelevance
	})

	if len(relevantCharacters) > maxRelevantCharacters {
		relevantCharacters = relevantCharacters[:maxRelevantCharacters]
	}

	// Quotes
	mainQuote := character.MainQuote()
	quotes := character.Quotes()

	arn.SortQuotesPopularFirst(quotes)

	// Set OpenGraph attributes
	description := utils.CutLongDescription(character.Description)

	customCtx := ctx.(*middleware.OpenGraphContext)
	customCtx.OpenGraph = &arn.OpenGraph{
		Tags: map[string]string{
			"og:title":       character.Name.Canonical,
			"og:image":       "https:" + character.ImageLink("large"),
			"og:url":         "https://" + assets.Domain + character.Link(),
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
		friendIDs := utils.Intersection(character.Likes, user.FollowIDs)
		friendObjects := arn.DB.GetMany("User", friendIDs)

		for _, obj := range friendObjects {
			if obj == nil {
				continue
			}

			friends = append(friends, obj.(*arn.User))
		}
	}

	return ctx.HTML(components.CharacterDetails(character, characterAnime, quotes, friends, relevantCharacters, mainQuote, user))
}
