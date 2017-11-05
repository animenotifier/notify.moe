package artworks

import (
	"github.com/aerogo/aero"
)

// Get artworks.
func Get(ctx *aero.Context) string {
	return ctx.HTML("Coming soon™.")
}
