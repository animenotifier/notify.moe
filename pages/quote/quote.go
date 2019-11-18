package quote

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/assets"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/server/middleware"
)

// Get quote.
func Get(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)
	id := ctx.Get("id")
	quote, err := arn.GetQuote(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Quote not found", err)
	}

	openGraph := &arn.OpenGraph{
		Tags: map[string]string{
			"og:title":       "Quote",
			"og:description": quote.Text.English,
			"og:url":         "https://" + assets.Domain + quote.Link(),
			"og:site_name":   "notify.moe",
			"og:type":        "article",
		},
	}

	character, _ := arn.GetCharacter(quote.CharacterID)

	if character != nil {
		openGraph.Tags["og:title"] = character.Name.Canonical + "'s quote"
		openGraph.Tags["og:image"] = "https:" + character.ImageLink("large")
	}

	customCtx := ctx.(*middleware.OpenGraphContext)
	customCtx.OpenGraph = openGraph
	return ctx.HTML(components.QuotePage(quote, character, user))
}
