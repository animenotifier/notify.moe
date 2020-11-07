package api

import (
	"net/http"
	"path"
	"sort"

	"github.com/aerogo/aero"
	"github.com/akyoto/color"
	"github.com/akyoto/uuid"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/arn/autodocs"
	"github.com/animenotifier/notify.moe/components"
)

// Get api page.
func Get(ctx aero.Context) error {
	types := []*autodocs.Type{}
	user := arn.GetUserFromContext(ctx)

	if user == nil {
		return ctx.Error(http.StatusUnauthorized, "Not logged in")
	}

	for typeName := range arn.DB.Types() {
		if typeName == "Session" {
			continue
		}

		typ, err := autodocs.GetTypeDocumentation(typeName, path.Join(arn.Root, "arn", typeName+".go"))
		types = append(types, typ)

		if err != nil {
			color.Red(err.Error())
			continue
		}
	}

	sort.Slice(types, func(i, j int) bool {
		return types[i].Name < types[j].Name
	})

	if (user.APIToken == uuid.UUID{}) {
		user.CreateAPIToken()
	}

	return ctx.HTML(components.API(types, user))
}
