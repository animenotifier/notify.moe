package activity

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
)

// Global activity page.
func Global(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)
	activities := fetchActivities(user, false)
	return render(ctx, activities)
}

// Followed activity page.
func Followed(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)
	activities := fetchActivities(user, true)
	return render(ctx, activities)
}

// fetchActivities filters the activities by the given filters.
func fetchActivities(user *arn.User, followedOnly bool) []arn.Activity {
	activities := arn.FilterActivityCreates(func(activity arn.Activity) bool {
		if followedOnly && user != nil && !user.IsFollowing(activity.GetCreatedBy()) {
			return false
		}

		if !activity.Creator().HasNick() {
			return false
		}

		obj := activity.(*arn.ActivityCreate).Object()

		if obj == nil {
			return false
		}

		draft, isDraftable := obj.(arn.Draftable)
		return !isDraftable || !draft.GetIsDraft()
	})

	arn.SortActivitiesLatestFirst(activities)
	return activities
}
