package activity

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils/infinitescroll"
)

const (
	activitiesFirstLoad = 40
	activitiesPerScroll = 20
)

// render renders the activities page with the given activities.
func render(ctx aero.Context, allActivities []arn.Activity) error {
	user := arn.GetUserFromContext(ctx)
	index, _ := ctx.GetInt("index")

	// Slice the part that we need
	activities := allActivities[index:]
	maxLength := activitiesFirstLoad

	if index > 0 {
		maxLength = activitiesPerScroll
	}

	if len(activities) > maxLength {
		activities = activities[:maxLength]
	}

	// Next index
	nextIndex := infinitescroll.NextIndex(ctx, len(allActivities), maxLength, index)

	// In case we're scrolling, send activities only (without the page frame)
	if index > 0 {
		return ctx.HTML(components.ActivitiesScrollable(activities, user))
	}

	// Otherwise, send the full page
	return ctx.HTML(components.ActivityFeed(activities, nextIndex, user))
}
