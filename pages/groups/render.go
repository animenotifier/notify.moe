package groups

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils/infinitescroll"
)

const (
	groupsFirstLoad = 27
	groupsPerScroll = 18
)

// render renders the groups page with the given groups.
func render(ctx aero.Context, allGroups []*arn.Group) error {
	user := arn.GetUserFromContext(ctx)
	index, _ := ctx.GetInt("index")

	// Slice the part that we need
	groups := allGroups[index:]
	maxLength := groupsFirstLoad

	if index > 0 {
		maxLength = groupsPerScroll
	}

	if len(groups) > maxLength {
		groups = groups[:maxLength]
	}

	// Next index
	nextIndex := infinitescroll.NextIndex(ctx, len(allGroups), maxLength, index)

	// In case we're scrolling, send groups only (without the page frame)
	if index > 0 {
		return ctx.HTML(components.GroupsScrollable(groups, user))
	}

	// Otherwise, send the full page
	return ctx.HTML(components.Groups(groups, nextIndex, user))
}
