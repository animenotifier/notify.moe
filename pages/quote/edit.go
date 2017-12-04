package quote

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
	"github.com/animenotifier/notify.moe/utils/editform"
)

// Edit track.
func Edit(ctx *aero.Context) string {
	user := utils.GetUser(ctx)
	id := ctx.Get("id")
	quote, err := arn.GetQuote(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Quote not found", err)
	}

	ctx.Data = &arn.OpenGraph{
		Tags: map[string]string{
			"og:title":     quote.Description,
			"og:url":       "https://" + ctx.App.Config.Domain + quote.Link(),
			"og:site_name": "notify.moe",
		},
	}

	return ctx.HTML(components.QuoteTabs(quote, user) + editform.Render(quote, "Edit quote", user))
}
