package company

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
	id := ctx.Get("id")
	company, err := arn.GetCompany(id)
	user := utils.GetUser(ctx)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Company not found", err)
	}

	ctx.Data = &arn.OpenGraph{
		Tags: map[string]string{
			"og:title":     company.Name.English,
			"og:url":       "https://" + ctx.App.Config.Domain + company.Link(),
			"og:site_name": "notify.moe",
			"og:image":     company.Image,
		},
	}

	return ctx.HTML(components.CompanyTabs(company, user) + editform.Render(company, "Edit company", user))
}
