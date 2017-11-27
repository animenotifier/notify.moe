package companies

import (
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

const maxEntries = 24

// Get renders the companies page.
func Get(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	companies := arn.FilterCompanies(func(company *arn.Company) bool {
		return !company.IsDraft
	})

	sort.Slice(companies, func(i, j int) bool {
		return companies[i].Created > companies[j].Created
	})

	if len(companies) > maxEntries {
		companies = companies[:maxEntries]
	}

	return ctx.HTML(components.Companies(companies, user))
}
