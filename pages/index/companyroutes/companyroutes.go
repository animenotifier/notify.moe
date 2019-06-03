package companyroutes

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/pages/companies"
	"github.com/animenotifier/notify.moe/pages/company"
	"github.com/animenotifier/notify.moe/utils/page"
)

// Register registers the page routes.
func Register(app *aero.Application) {
	// Companies
	page.Get(app, "/company/:id", company.Get)
	page.Get(app, "/company/:id/edit", company.Edit)
	page.Get(app, "/company/:id/history", company.History)
	page.Get(app, "/companies", companies.Popular)
	page.Get(app, "/companies/from/:index", companies.Popular)
	page.Get(app, "/companies/all", companies.All)
}
