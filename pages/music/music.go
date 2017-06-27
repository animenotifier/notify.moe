package music

import "github.com/aerogo/aero"

// Get renders the music page.
func Get(ctx *aero.Context) string {
	return ctx.HTML("Coming soon.")
}
