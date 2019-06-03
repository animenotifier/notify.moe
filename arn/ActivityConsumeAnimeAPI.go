package arn

import (
	"errors"

	"github.com/aerogo/aero"
	"github.com/aerogo/api"
)

// Force interface implementations
var (
	_ Activity      = (*ActivityConsumeAnime)(nil)
	_ api.Deletable = (*ActivityConsumeAnime)(nil)
	_ api.Savable   = (*ActivityConsumeAnime)(nil)
)

// Authorize returns an error if the given API POST request is not authorized.
func (activity *ActivityConsumeAnime) Authorize(ctx aero.Context, action string) error {
	user := GetUserFromContext(ctx)

	if user == nil {
		return errors.New("Not logged in")
	}

	if user.ID != activity.CreatedBy {
		return errors.New("Can't modify activities from other users")
	}

	return nil
}

// Save saves the activity object in the database.
func (activity *ActivityConsumeAnime) Save() {
	DB.Set("ActivityConsumeAnime", activity.ID, activity)
}

// DeleteInContext deletes the activity in the given context.
func (activity *ActivityConsumeAnime) DeleteInContext(ctx aero.Context) error {
	return activity.Delete()
}

// Delete deletes the object from the database.
func (activity *ActivityConsumeAnime) Delete() error {
	DB.Delete("ActivityConsumeAnime", activity.ID)
	return nil
}

// // Force interface implementations
// var (
// 	_ Likeable          = (*Activity)(nil)
// 	_ LikeEventReceiver = (*Activity)(nil)
// 	_ api.Deletable     = (*Activity)(nil)
// )

// // Actions
// func init() {
// 	API.RegisterActions("Activity", []*api.Action{
// 		// Like
// 		LikeAction(),

// 		// Unlike
// 		UnlikeAction(),
// 	})
// }

// // Authorize returns an error if the given API request is not authorized.
// func (activity *Activity) Authorize(ctx aero.Context, action string) error {
// 	user := GetUserFromContext(ctx)

// 	if user == nil {
// 		return errors.New("Not logged in")
// 	}

// 	return nil
// }

// // DeleteInContext deletes the activity in the given context.
// func (activity *Activity) DeleteInContext(ctx aero.Context) error {
// 	return activity.Delete()
// }
