package post

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// NewPostArea renders a new post area.
func NewPostArea(ctx *aero.Context) string {
	user := utils.GetUser(ctx)
	return ctx.HTML(components.NewPostArea(user, "Reply"))
}
