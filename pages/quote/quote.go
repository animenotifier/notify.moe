package quote

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Get quote.
func Get(ctx *aero.Context) string {
	user := utils.GetUser(ctx)
	id := ctx.Get("id")
	quote, err := arn.GetQuote(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Quote not found", err)
	}

	openGraph := &arn.OpenGraph{
		Tags: map[string]string{
			"og:title":       "Quote",
			"og:description": quote.Text.English,
			"og:url":         "https://" + ctx.App.Config.Domain + quote.Link(),
			"og:site_name":   "notify.moe",
			"og:type":        "article",
		},
	}

	character, _ := arn.GetCharacter(quote.CharacterID)

	if character != nil {
		openGraph.Tags["og:title"] = character.Name.Canonical + "'s quote"
		openGraph.Tags["og:image"] = "https:" + character.ImageLink("large")
	}

	ctx.Data = openGraph
	return ctx.HTML(components.QuotePage(quote, character, user))
}
