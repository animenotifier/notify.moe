package animeepisode

import "github.com/aerogo/aero"

func Get(ctx *aero.Context) string {
	return ctx.HTML("")
}
