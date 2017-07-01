package apiview

import (
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

// Get api page.
func Get(ctx *aero.Context) string {
	types := []string{}

	for typeName := range arn.DB.Types() {
		types = append(types, typeName)
	}

	sort.Strings(types)

	return ctx.HTML(components.API(types))
}
