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

	character, err := arn.GetCharacter(quote.CharacterID)
	if err != nil {
		return ctx.Error(http.StatusNotFound, "Quote not found", err)
	}

	openGraph := &arn.OpenGraph{
		Tags: map[string]string{
			"og:title":       quote.Description,
			"og:description": quote.Description,
			"og:url":         "https://" + ctx.App.Config.Domain + quote.Link(),
			"og:site_name":   "notify.moe",
			"og:type":        "article",
		},
	}

	ctx.Data = openGraph
	return ctx.HTML(components.QuotePage(quote, character, user))
}
