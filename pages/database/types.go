package database

import (
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
)

// Types shows which types are available in the database.
func Types(ctx aero.Context) error {
	typeMap := arn.DB.Types()
	types := make([]string, 0, len(typeMap))

	for typeName := range typeMap {
		if arn.IsPrivateType(typeName) {
			continue
		}

		types = append(types, typeName)
	}

	sort.Strings(types)
	return ctx.JSON(types)
}
