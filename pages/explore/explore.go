package explore

import (
	"github.com/aerogo/aero"
)

// Get ...
func Get(ctx *aero.Context) string {
	// var cache arn.ListOfIDs
	// err := arn.DB.GetObject("Cache", "airing anime", &cache)

	// airing, err := arn.GetAiringAnimeCached()

	// if err != nil {
	// 	return ctx.Error(500, "Couldn't fetch airing anime", err)
	// }

	// return ctx.HTML(components.Airing(airing))
	return ctx.HTML("Not implemented")
}
