package company

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/assets"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/server/middleware"
	"github.com/animenotifier/notify.moe/utils/editform"
)

// Edit company.
func Edit(ctx aero.Context) error {
	id := ctx.Get("id")
	company, err := arn.GetCompany(id)
	user := arn.GetUserFromContext(ctx)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Company not found", err)
	}

	customCtx := ctx.(*middleware.OpenGraphContext)
	customCtx.OpenGraph = &arn.OpenGraph{
		Tags: map[string]string{
			"og:title":     company.Name.English,
			"og:url":       "https://" + assets.Domain + company.Link(),
			"og:site_name": "notify.moe",
			// "og:image":     company.Image,
		},
	}

	return ctx.HTML(components.CompanyTabs(company, user) + editform.Render(company, "Edit company", user))
}
