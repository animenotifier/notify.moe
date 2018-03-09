package database

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/components"
)

// Get the dashboard.
func Get(ctx *aero.Context) string {
	return ctx.HTML(components.Database(ctx.URI()))
}
