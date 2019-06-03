package arn

import (
	"errors"
	"reflect"

	"github.com/aerogo/aero"
	"github.com/aerogo/api"
)

// Likeable ...
type Likeable interface {
	Like(userID UserID)
	Unlike(userID UserID)
	LikedBy(userID UserID) bool
	CountLikes() int
	Link() string
	Save()
}

// LikeEventReceiver ...
type LikeEventReceiver interface {
	OnLike(user *User)
}

// LikeAction ...
func LikeAction() *api.Action {
	return &api.Action{
		Name:  "like",
		Route: "/like",
		Run: func(obj interface{}, ctx aero.Context) error {
			field := reflect.ValueOf(obj).Elem().FieldByName("IsDraft")

			if field.IsValid() && field.Bool() {
				return errors.New("Drafts need to be published before they can be liked")
			}

			likeable := obj.(Likeable)
			user := GetUserFromContext(ctx)

			if user == nil {
				return errors.New("Not logged in")
			}

			likeable.Like(user.ID)

			// Call OnLike if the object implements it
			receiver, ok := likeable.(LikeEventReceiver)

			if ok {
				receiver.OnLike(user)
			}

			likeable.Save()
			return nil
		},
	}
}

// UnlikeAction ...
func UnlikeAction() *api.Action {
	return &api.Action{
		Name:  "unlike",
		Route: "/unlike",
		Run: func(obj interface{}, ctx aero.Context) error {
			likeable := obj.(Likeable)
			user := GetUserFromContext(ctx)

			if user == nil {
				return errors.New("Not logged in")
			}

			likeable.Unlike(user.ID)
			likeable.Save()
			return nil
		},
	}
}
