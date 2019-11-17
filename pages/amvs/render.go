package amvs

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils/infinitescroll"
)

const (
	amvsFirstLoad = 12
	amvsPerScroll = 9
)

// render renders the AMVs page with the given AMVs.
func render(ctx aero.Context, allAMVs []*arn.AMV) error {
	user := arn.GetUserFromContext(ctx)
	index, _ := ctx.GetInt("index")
	tag := ctx.Get("tag")

	// Slice the part that we need
	amvs := allAMVs[index:]
	maxLength := amvsFirstLoad

	if index > 0 {
		maxLength = amvsPerScroll
	}

	if len(amvs) > maxLength {
		amvs = amvs[:maxLength]
	}

	// Next index
	nextIndex := infinitescroll.NextIndex(ctx, len(allAMVs), maxLength, index)

	// In case we're scrolling, send AMVs only (without the page frame)
	if index > 0 {
		return ctx.HTML(components.AMVsScrollable(amvs, user))
	}

	// Otherwise, send the full page
	return ctx.HTML(components.AMVs(amvs, nextIndex, tag, user))
}
