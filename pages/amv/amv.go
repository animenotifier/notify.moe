package amv

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/middleware"
	"github.com/animenotifier/notify.moe/utils"
)

// Get a single AMV.
func Get(ctx aero.Context) error {
	id := ctx.Get("id")
	amv, err := arn.GetAMV(id)
	user := utils.GetUser(ctx)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "AMV not found", err)
	}

	customCtx := ctx.(*middleware.OpenGraphContext)
	customCtx.OpenGraph = getOpenGraph(ctx, amv)
	return ctx.HTML(components.AMVPage(amv, user))
}
