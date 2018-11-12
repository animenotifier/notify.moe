package activity

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/utils"
)

// Global activity page.
func Global(ctx *aero.Context) string {
	user := utils.GetUser(ctx)
	activities := fetchActivities(user, false)
	return render(ctx, activities)
}

// Followed activity page.
func Followed(ctx *aero.Context) string {
	user := utils.GetUser(ctx)
	activities := fetchActivities(user, true)
	return render(ctx, activities)
}

// fetchActivities filters the activities by the given filters.
func fetchActivities(user *arn.User, followedOnly bool) []arn.Activity {
	var followedUserIDs []string

	if followedOnly && user != nil {
		followedUserIDs = user.Follows().Items
	}

	activities := arn.FilterActivities(func(activity arn.Activity) bool {
		if followedOnly && !arn.Contains(followedUserIDs, activity.GetCreatedBy()) {
			return false
		}

		if activity.Type() == "ActivityCreate" {
			obj := activity.(*arn.ActivityCreate).Object()

			if obj == nil {
				return false
			}

			draft, isDraftable := obj.(arn.Draftable)
			return !isDraftable || !draft.GetIsDraft()
		}

		if activity.Type() == "ActivityConsumeAnime" {
			return activity.(*arn.ActivityConsumeAnime).Anime() != nil
		}

		return false
	})

	arn.SortActivitiesLatestFirst(activities)
	return activities
}
