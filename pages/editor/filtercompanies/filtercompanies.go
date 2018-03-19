package filtercompanies

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

const maxEntries = 70

// NoDescription ...
func NoDescription(ctx *aero.Context) string {
	companies := arn.FilterCompanies(func(company *arn.Company) bool {
		return !company.IsDraft && len(company.Description) < 5
	})

	arn.SortCompaniesPopularFirst(companies)

	count := len(companies)

	if count > maxEntries {
		companies = companies[:maxEntries]
	}

	return ctx.HTML(components.CompaniesEditorList(companies, count, ctx.URI()))
}
