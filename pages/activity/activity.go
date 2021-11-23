package activity

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
)

// Posts activity page.
func Posts(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)
	activities := fetchCreateActivities(user)
	return render(ctx, activities)
}

// Watch activity page.
func Watch(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)
	activities := fetchConsumeActivities(user)
	return render(ctx, activities)
}

// fetchCreateActivities filters the activities by the given filters.
func fetchCreateActivities(user *arn.User) []arn.Activity {
	followedOnly := user != nil && user.Settings().Activity.ShowFollowedOnly

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

// fetchConsumeActivities filters the consume activities by the given filters.
func fetchConsumeActivities(user *arn.User) []arn.Activity {
	followedOnly := user != nil && user.Settings().Activity.ShowFollowedOnly

	activities := arn.FilterActivitiesConsumeAnime(func(activity arn.Activity) bool {
		if followedOnly && user != nil && !user.IsFollowing(activity.GetCreatedBy()) {
			return false
		}

		if !activity.Creator().HasNick() {
			return false
		}

		return true
	})

	arn.SortActivitiesLatestFirst(activities)
	return activities
}
