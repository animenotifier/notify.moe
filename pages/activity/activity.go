package activity

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

const maxActivitiesPerPage = 40

// Get activity page.
func Get(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	activities := arn.FilterActivities(func(activity arn.Activity) bool {
		if activity.Type() == "ActivityCreate" {
			obj := activity.(*arn.ActivityCreate).Object()

			if obj == nil {
				return false
			}

			draft, isDraftable := obj.(arn.HasDraft)
			return !isDraftable || !draft.IsDraft
		}

		if activity.Type() == "ActivityConsumeAnime" {
			return activity.(*arn.ActivityConsumeAnime).Anime() != nil
		}

		return false
	})

	arn.SortActivitiesLatestFirst(activities)

	if len(activities) > maxActivitiesPerPage {
		activities = activities[:maxActivitiesPerPage]
	}

	return ctx.HTML(components.ActivityFeed(activities, user))
}
