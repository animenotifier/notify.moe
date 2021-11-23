package arn

import (
	"errors"

	"github.com/aerogo/aero"
	"github.com/aerogo/api"
)

// Add adds an user to the user if it hasn't been added yet.
func (user *User) Follow(followUserID UserID) error {
	if followUserID == user.ID {
		return errors.New("You can't follow yourself")
	}

	if user.IsFollowing(followUserID) {
		return errors.New("User " + followUserID + " has already been added")
	}

	followedUser, err := GetUser(followUserID)

	if err != nil {
		return err
	}

	user.FollowIDs = append(user.FollowIDs, followUserID)

	// Send notification
	if !followedUser.Settings().Notification.NewFollowers {
		return nil
	}

	followedUser.SendNotification(&PushNotification{
		Title:   "You have a new follower!",
		Message: user.Nick + " started following you.",
		Icon:    "https:" + user.AvatarLink("large"),
		Link:    "https://notify.moe" + user.Link(),
		Type:    NotificationTypeFollow,
	})

	return nil
}

// Unfollow removes the user ID from the follow list.
func (user *User) Unfollow(userID UserID) bool {
	for index, item := range user.FollowIDs {
		if item == userID {
			user.FollowIDs = append(user.FollowIDs[:index], user.FollowIDs[index+1:]...)
			return true
		}
	}

	return false
}

// IsFollowing checks if the object follows the user ID.
func (user *User) IsFollowing(userID UserID) bool {
	if userID == user.ID {
		return true
	}

	for _, item := range user.FollowIDs {
		if item == userID {
			return true
		}
	}

	return false
}

// Follows returns a slice of all the users you are following.
func (user *User) Follows() []*User {
	followsObj := DB.GetMany("User", user.FollowIDs)
	follows := make([]*User, len(followsObj))

	for i, user := range followsObj {
		follows[i] = user.(*User)
	}

	return follows
}

// Friends returns a slice of all the users you are following that also follow you.
func (user *User) Friends() []*User {
	followsObj := DB.GetMany("User", user.FollowIDs)
	friends := make([]*User, 0, len(followsObj))

	for _, friendObj := range followsObj {
		friend := friendObj.(*User)

		if friend.IsFollowing(user.ID) {
			friends = append(friends, friend)
		}
	}

	return friends
}

// Followers returns the users who follow the user.
func (user *User) Followers() []*User {
	var followerIDs []string

	for follower := range StreamUsers() {
		if follower.IsFollowing(user.ID) {
			followerIDs = append(followerIDs, follower.ID)
		}
	}

	usersObj := DB.GetMany("User", followerIDs)
	users := make([]*User, len(usersObj))

	for i, obj := range usersObj {
		users[i] = obj.(*User)
	}

	return users
}

// FollowersCount returns how many followers the user has.
func (user *User) FollowersCount() int {
	count := 0

	for follower := range StreamUsers() {
		if follower.IsFollowing(user.ID) {
			count++
		}
	}

	return count
}

// UserFollowerCountMap returns a map of user ID keys and their corresping number of followers as the value.
func UserFollowerCountMap() map[string]int {
	followCount := map[string]int{}

	for user := range StreamUsers() {
		for _, followUserID := range user.FollowIDs {
			followCount[followUserID]++
		}
	}

	return followCount
}

// FollowAction returns an API action that adds a user ID to the follow list.
func FollowAction() *api.Action {
	return &api.Action{
		Name:  "follow",
		Route: "/follow/:follow-id",
		Run: func(obj interface{}, ctx aero.Context) error {
			user := obj.(*User)
			followID := ctx.Get("follow-id")
			err := user.Follow(followID)

			if err != nil {
				return err
			}

			user.Save()
			return nil
		},
	}
}

// UnfollowAction returns an API action that removes a user ID from the follow list.
func UnfollowAction() *api.Action {
	return &api.Action{
		Name:  "unfollow",
		Route: "/unfollow/:unfollow-id",
		Run: func(obj interface{}, ctx aero.Context) error {
			user := obj.(*User)
			unfollowID := ctx.Get("unfollow-id")

			if !user.Unfollow(unfollowID) {
				return errors.New("This item does not exist in the list")
			}

			user.Save()
			return nil
		},
	}
}
