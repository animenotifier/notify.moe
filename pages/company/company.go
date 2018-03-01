package company

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Get company.
func Get(ctx *aero.Context) string {
	user := utils.GetUser(ctx)
	id := ctx.Get("id")
	company, err := arn.GetCompany(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Company not found", err)
	}

	openGraph := &arn.OpenGraph{
		Tags: map[string]string{
			"og:title":       company.Name.English,
			"og:description": company.Description,
			"og:url":         "https://" + ctx.App.Config.Domain + company.Link(),
			"og:site_name":   "notify.moe",
			"og:type":        "article",
		},
	}

	if company.Image != "" {
		openGraph.Tags["og:image"] = company.Image
	}

	if company.Description != "" {
		openGraph.Tags["og:description"] = company.Description
	} else {
		openGraph.Tags["og:description"] = company.Name.English + " company information."
	}

	ctx.Data = openGraph

	studioAnime, producedAnime, licensedAnime := company.Anime()

	return ctx.HTML(components.CompanyPage(company, studioAnime, producedAnime, licensedAnime, user))
}
