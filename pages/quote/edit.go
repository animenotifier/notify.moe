package quote

import (
	"net/http"

	"github.com/animenotifier/notify.moe/server/middleware"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/assets"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils/editform"
)

// Edit quote.
func Edit(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)
	id := ctx.Get("id")
	quote, err := arn.GetQuote(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Quote not found", err)
	}

	customCtx := ctx.(*middleware.OpenGraphContext)
	customCtx.OpenGraph = &arn.OpenGraph{
		Tags: map[string]string{
			"og:title":     quote.Text.English,
			"og:url":       "https://" + assets.Domain + quote.Link(),
			"og:site_name": "notify.moe",
		},
	}

	if quote.Character() != nil {
		customCtx.OpenGraph.Tags["og:image"] = quote.Character().ImageLink("large")
	}

	return ctx.HTML(components.QuoteTabs(quote, user) + editform.Render(quote, "Edit quote", user))
}
