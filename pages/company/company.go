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

	var overallRating, storyRating, visualsRating, soundtrackRating float64
	var overallCount, storyCount, visualsCount, soundtrackCount int

	for _, anime := range studioAnime {
		overallRating += anime.Rating.Overall
		storyRating += anime.Rating.Story
		visualsRating += anime.Rating.Visuals
		soundtrackRating += anime.Rating.Soundtrack

		overallCount += anime.Rating.Count.Overall
		storyCount += anime.Rating.Count.Story
		visualsCount += anime.Rating.Count.Visuals
		soundtrackCount += anime.Rating.Count.Soundtrack
	}

	totalStudioAnime := float64(len(studioAnime))

	overallRating /= totalStudioAnime
	storyRating /= totalStudioAnime
	visualsRating /= totalStudioAnime
	soundtrackRating /= totalStudioAnime

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

	return ctx.HTML(components.CompanyPage(company, studioAnime, producedAnime, licensedAnime, closeCompanies, distances, user, overallRating, storyRating, visualsRating, soundtrackRating, overallCount, storyCount, visualsCount, soundtrackCount))
}
