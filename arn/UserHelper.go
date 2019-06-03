package arn

import (
	"errors"
	"sort"

	"github.com/aerogo/nano"
)

// GetUser fetches the user with the given ID from the database.
func GetUser(id UserID) (*User, error) {
	obj, err := DB.Get("User", id)

	if err != nil {
		return nil, err
	}

	return obj.(*User), nil
}

// GetUserByNick fetches the user with the given nick from the database.
func GetUserByNick(nick string) (*User, error) {
	obj, err := DB.Get("NickToUser", nick)

	if err != nil {
		return nil, err
	}

	userID := obj.(*NickToUser).UserID
	user, err := GetUser(userID)

	return user, err
}

// GetUserByEmail fetches the user with the given email from the database.
func GetUserByEmail(email string) (*User, error) {
	if email == "" {
		return nil, errors.New("Email is empty")
	}

	obj, err := DB.Get("EmailToUser", email)

	if err != nil {
		return nil, err
	}

	userID := obj.(*EmailToUser).UserID
	user, err := GetUser(userID)

	return user, err
}

// GetUserByFacebookID fetches the user with the given Facebook ID from the database.
func GetUserByFacebookID(facebookID string) (*User, error) {
	obj, err := DB.Get("FacebookToUser", facebookID)

	if err != nil {
		return nil, err
	}

	userID := obj.(*FacebookToUser).UserID
	user, err := GetUser(userID)

	return user, err
}

// GetUserByTwitterID fetches the user with the given Twitter ID from the database.
func GetUserByTwitterID(twitterID string) (*User, error) {
	obj, err := DB.Get("TwitterToUser", twitterID)

	if err != nil {
		return nil, err
	}

	userID := obj.(*TwitterToUser).UserID
	user, err := GetUser(userID)

	return user, err
}

// GetUserByGoogleID fetches the user with the given Google ID from the database.
func GetUserByGoogleID(googleID string) (*User, error) {
	obj, err := DB.Get("GoogleToUser", googleID)

	if err != nil {
		return nil, err
	}

	userID := obj.(*GoogleToUser).UserID
	user, err := GetUser(userID)

	return user, err
}

// StreamUsers returns a stream of all users.
func StreamUsers() <-chan *User {
	channel := make(chan *User, nano.ChannelBufferSize)

	go func() {
		for obj := range DB.All("User") {
			channel <- obj.(*User)
		}

		close(channel)
	}()

	return channel
}

// AllUsers returns a slice of all users.
func AllUsers() ([]*User, error) {
	all := make([]*User, 0, DB.Collection("User").Count())

	for obj := range StreamUsers() {
		all = append(all, obj)
	}

	return all, nil
}

// FilterUsers filters all users by a custom function.
func FilterUsers(filter func(*User) bool) []*User {
	var filtered []*User

	for obj := range StreamUsers() {
		if filter(obj) {
			filtered = append(filtered, obj)
		}
	}

	return filtered
}

// SortUsersLastSeenFirst sorts a list of users by their last seen date.
func SortUsersLastSeenFirst(users []*User) {
	sort.Slice(users, func(i, j int) bool {
		return users[i].LastSeen > users[j].LastSeen
	})
}

// SortUsersLastSeenLast sorts a list of users by their last seen date.
func SortUsersLastSeenLast(users []*User) {
	sort.Slice(users, func(i, j int) bool {
		return users[i].LastSeen < users[j].LastSeen
	})
}

// SortUsersFollowers sorts a list of users by their number of followers.
func SortUsersFollowers(users []*User) {
	followCount := UserFollowerCountMap()

	sort.Slice(users, func(i, j int) bool {
		if users[i].HasAvatar() != users[j].HasAvatar() {
			return users[i].HasAvatar()
		}

		followersA := followCount[users[i].ID]
		followersB := followCount[users[j].ID]

		if followersA == followersB {
			return users[i].Nick < users[j].Nick
		}

		return followersA > followersB
	})
}
