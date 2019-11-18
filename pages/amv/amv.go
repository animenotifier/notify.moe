package amv

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/server/middleware"
)

// Get a single AMV.
func Get(ctx aero.Context) error {
	id := ctx.Get("id")
	amv, err := arn.GetAMV(id)
	user := arn.GetUserFromContext(ctx)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "AMV not found", err)
	}

	customCtx := ctx.(*middleware.OpenGraphContext)
	customCtx.OpenGraph = getOpenGraph(amv)
	return ctx.HTML(components.AMVPage(amv, user))
}
