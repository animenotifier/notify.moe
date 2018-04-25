package companyroutes

import (
	"github.com/aerogo/layout"
	"github.com/animenotifier/notify.moe/pages/companies"
	"github.com/animenotifier/notify.moe/pages/company"
)

// Register registers the page routes.
func Register(l *layout.Layout) {
	// Companies
	l.Page("/company/:id", company.Get)
	l.Page("/company/:id/edit", company.Edit)
	l.Page("/company/:id/history", company.History)
	l.Page("/companies", companies.Popular)
	l.Page("/companies/from/:index", companies.Popular)
	l.Page("/companies/all", companies.All)
}
