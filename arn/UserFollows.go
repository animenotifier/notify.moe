package arn

import (
	"errors"

	"github.com/aerogo/nano"
)

// UserFollows is a list including IDs to users you follow.
type UserFollows struct {
	UserID UserID   `json:"userId"`
	Items  []string `json:"items"`
}

// NewUserFollows creates a new UserFollows list.
func NewUserFollows(userID UserID) *UserFollows {
	return &UserFollows{
		UserID: userID,
		Items:  []string{},
	}
}

// Add adds an user to the list if it hasn't been added yet.
func (list *UserFollows) Add(userID UserID) error {
	if userID == list.UserID {
		return errors.New("You can't follow yourself")
	}

	if list.Contains(userID) {
		return errors.New("User " + userID + " has already been added")
	}

	list.Items = append(list.Items, userID)

	// Send notification
	user, err := GetUser(userID)

	if err == nil {
		if !user.Settings().Notification.NewFollowers {
			return nil
		}

		follower, err := GetUser(list.UserID)

		if err == nil {
			user.SendNotification(&PushNotification{
				Title:   "You have a new follower!",
				Message: follower.Nick + " started following you.",
				Icon:    "https:" + follower.AvatarLink("large"),
				Link:    "https://notify.moe" + follower.Link(),
				Type:    NotificationTypeFollow,
			})
		}
	}

	return nil
}

// Remove removes the user ID from the list.
func (list *UserFollows) Remove(userID UserID) bool {
	for index, item := range list.Items {
		if item == userID {
			list.Items = append(list.Items[:index], list.Items[index+1:]...)
			return true
		}
	}

	return false
}

// Contains checks if the list contains the user ID already.
func (list *UserFollows) Contains(userID UserID) bool {
	for _, item := range list.Items {
		if item == userID {
			return true
		}
	}

	return false
}

// Users returns a slice of all the users you are following.
func (list *UserFollows) Users() []*User {
	followsObj := DB.GetMany("User", list.Items)
	follows := make([]*User, len(followsObj))

	for i, obj := range followsObj {
		follows[i] = obj.(*User)
	}

	return follows
}

// UsersWhoFollowBack returns a slice of all the users you are following that also follow you.
func (list *UserFollows) UsersWhoFollowBack() []*User {
	followsObj := DB.GetMany("User", list.Items)
	friends := make([]*User, 0, len(followsObj))

	for _, obj := range followsObj {
		friend := obj.(*User)

		if Contains(friend.Follows().Items, list.UserID) {
			friends = append(friends, friend)
		}
	}

	return friends
}

// UserFollowerCountMap returns a map of user ID keys and their corresping number of followers as the value.
func UserFollowerCountMap() map[string]int {
	followCount := map[string]int{}

	for list := range StreamUserFollows() {
		for _, followUserID := range list.Items {
			followCount[followUserID]++
		}
	}

	return followCount
}

// GetUserFollows ...
func GetUserFollows(id UserID) (*UserFollows, error) {
	obj, err := DB.Get("UserFollows", id)

	if err != nil {
		return nil, err
	}

	return obj.(*UserFollows), nil
}

// StreamUserFollows returns a stream of all user follows.
func StreamUserFollows() <-chan *UserFollows {
	channel := make(chan *UserFollows, nano.ChannelBufferSize)

	go func() {
		for obj := range DB.All("UserFollows") {
			channel <- obj.(*UserFollows)
		}

		close(channel)
	}()

	return channel
}

// AllUserFollows returns a slice of all user follows.
func AllUserFollows() ([]*UserFollows, error) {
	all := make([]*UserFollows, 0, DB.Collection("UserFollows").Count())

	for obj := range StreamUserFollows() {
		all = append(all, obj)
	}

	return all, nil
}
