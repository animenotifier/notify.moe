package amvs

import (
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Latest AMVs.
func Latest(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	amvs := arn.FilterAMVs(func(amv *arn.AMV) bool {
		return !amv.IsDraft
	})

	sort.Slice(amvs, func(i, j int) bool {
		return amvs[i].Created > amvs[j].Created
	})

	return ctx.HTML(components.AMVs(amvs, -1, "", user))
}

// Best AMVs.
func Best(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	amvs := arn.FilterAMVs(func(amv *arn.AMV) bool {
		return !amv.IsDraft
	})

	sort.Slice(amvs, func(i, j int) bool {
		if len(amvs[i].Likes) == len(amvs[j].Likes) {
			return amvs[i].Title.String() < amvs[j].Title.String()
		}

		return len(amvs[i].Likes) > len(amvs[j].Likes)
	})

	return ctx.HTML(components.AMVs(amvs, -1, "", user))
}
