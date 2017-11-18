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

	return ctx.HTML(components.CompanyPage(company, user))
}
