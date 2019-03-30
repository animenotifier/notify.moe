package apiview

import (
	"path"
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/arn/autodocs"
	"github.com/animenotifier/notify.moe/components"
	"github.com/blitzprog/color"
)

// Get api page.
func Get(ctx *aero.Context) string {
	types := []*autodocs.Type{}

	for typeName := range arn.DB.Types() {
		if typeName == "Session" {
			continue
		}

		typ, err := autodocs.GetTypeDocumentation(typeName, path.Join(arn.Root, "..", "arn", typeName+".go"))
		types = append(types, typ)

		if err != nil {
			color.Red(err.Error())
			continue
		}
	}

	sort.Slice(types, func(i, j int) bool {
		return types[i].Name < types[j].Name
	})

	return ctx.HTML(components.API(types))
}
