package admin

import (
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Get admin page.
func Get(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	if user == nil || user.Role != "admin" {
		return ctx.Redirect("/")
	}

	types := []string{}

	for typeName := range arn.DB.Types() {
		types = append(types, typeName)
	}

	sort.Strings(types)

	return ctx.HTML(components.Admin(user, types))
}
