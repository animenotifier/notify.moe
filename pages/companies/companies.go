package companies

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

const maxEntries = 12

// Get renders the companies page.
func Get(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	companies := arn.FilterCompanies(func(company *arn.Company) bool {
		return !company.IsDraft
	})

	if len(companies) > maxEntries {
		companies = companies[:maxEntries]
	}

	return ctx.HTML(components.Companies(companies, user))
}
