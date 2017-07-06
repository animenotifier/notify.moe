package layout

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Render layout.
func Render(ctx *aero.Context, content string) string {
	user := utils.GetUser(ctx)
	meta, _ := ctx.Data.(map[string]string)
	return components.Layout(ctx.App, ctx, user, meta, content)
}
