package amvs

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Latest AMVs.
func Latest(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	return ctx.HTML(components.AMVs(nil, -1, "", user))
}

// Best AMVs.
func Best(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	return ctx.HTML(components.AMVs(nil, -1, "", user))
}
