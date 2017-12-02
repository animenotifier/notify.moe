package companies

import (
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

const maxPopularCompanies = 100

// Popular renders the companies sorted by popularity.
func Popular(ctx *aero.Context) string {
	user := utils.GetUser(ctx)
	companies := []*arn.Company{}

	// ID to popularity
	popularity := map[string]int{}

	for anime := range arn.StreamAnime() {
		for _, studio := range anime.Studios() {
			popularity[studio.ID] += anime.Popularity.Watching + anime.Popularity.Completed
		}
	}

	for companyID := range popularity {
		company, err := arn.GetCompany(companyID)

		if err != nil {
			continue
		}

		companies = append(companies, company)
	}

	sort.Slice(companies, func(i, j int) bool {
		a := companies[i]
		b := companies[j]

		aPopularity := popularity[a.ID]
		bPopularity := popularity[b.ID]

		if aPopularity == bPopularity {
			return a.Name.English < b.Name.English
		}

		return aPopularity > bPopularity
	})

	if len(companies) > maxPopularCompanies {
		companies = companies[:maxPopularCompanies]
	}

	return ctx.HTML(components.PopularCompanies(companies, popularity, user))
}
