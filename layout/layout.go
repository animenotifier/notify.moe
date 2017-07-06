package layout

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Render layout.
func Render(ctx *aero.Context, content string) string {
	user := utils.GetUser(ctx)
	openGraph, _ := ctx.Data.(*arn.OpenGraph)
	return components.Layout(ctx.App, ctx, user, openGraph, content)
}
