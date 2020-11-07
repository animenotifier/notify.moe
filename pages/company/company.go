package company

import (
	"net/http"
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/assets"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/server/middleware"
	"github.com/animenotifier/notify.moe/utils"
)

// Get renders a company page.
func Get(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)
	id := ctx.Get("id")
	company, err := arn.GetCompany(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Company not found", err)
	}

	description := utils.CutLongDescription(company.Description)

	openGraph := &arn.OpenGraph{
		Tags: map[string]string{
			"og:title":     company.Name.English,
			"og:url":       "https://" + assets.Domain + company.Link(),
			"og:site_name": "notify.moe",
			"og:type":      "article",
		},
	}

	// if company.Image != "" {
	// 	openGraph.Tags["og:image"] = company.Image
	// }

	if description != "" {
		openGraph.Tags["og:description"] = description
	} else {
		openGraph.Tags["og:description"] = company.Name.English + " company information."
	}

	customCtx := ctx.(*middleware.OpenGraphContext)
	customCtx.OpenGraph = openGraph

	studioAnime, producedAnime, licensedAnime := company.Anime()

	// Find close companies
	var closeCompanies []*arn.Company
	distances := map[string]float64{}

	// Average rating
	rating := &arn.AnimeRating{}

	for _, anime := range studioAnime {
		rating.Overall += anime.Rating.Overall
		rating.Story += anime.Rating.Story
		rating.Visuals += anime.Rating.Visuals
		rating.Soundtrack += anime.Rating.Soundtrack

		rating.Count.Overall += anime.Rating.Count.Overall
		rating.Count.Story += anime.Rating.Count.Story
		rating.Count.Visuals += anime.Rating.Count.Visuals
		rating.Count.Soundtrack += anime.Rating.Count.Soundtrack
	}

	totalStudioAnime := float64(len(studioAnime))

	rating.Overall /= totalStudioAnime
	rating.Story /= totalStudioAnime
	rating.Visuals /= totalStudioAnime
	rating.Soundtrack /= totalStudioAnime

	if company.Location.IsValid() {
		closeCompanies = arn.FilterCompanies(func(closeCompany *arn.Company) bool {
			if closeCompany.ID == company.ID {
				return false
			}

			if !closeCompany.Location.IsValid() {
				return false
			}

			distance := company.Location.Distance(closeCompany.Location)
			distances[closeCompany.ID] = distance

			return distance <= 1.0
		})

		sort.Slice(closeCompanies, func(i, j int) bool {
			return distances[closeCompanies[i].ID] < distances[closeCompanies[j].ID]
		})
	}

	return ctx.HTML(components.CompanyPage(company, studioAnime, producedAnime, licensedAnime, closeCompanies, distances, rating, user))
}
