package airing

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

// Get ...
func Get(ctx *aero.Context) string {
	var cache arn.ListOfIDs
	err := arn.GetObject("Cache", "airing anime", &cache)

	airing, err := arn.GetAiringAnimeCached()

	if err != nil {
		return ctx.Error(500, "Couldn't fetch airing anime", err)
	}

	return ctx.HTML(components.Airing(airing))
}
