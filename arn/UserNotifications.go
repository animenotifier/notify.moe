package arn

import (
	"errors"

	"github.com/aerogo/nano"
)

// UserNotifications is a list including IDs to your notifications.
type UserNotifications struct {
	UserID UserID   `json:"userId"`
	Items  []string `json:"items"`
}

// NewUserNotifications creates a new UserNotifications list.
func NewUserNotifications(userID UserID) *UserNotifications {
	return &UserNotifications{
		UserID: userID,
		Items:  []string{},
	}
}

// CountUnseen returns the number of unseen notifications.
func (list *UserNotifications) CountUnseen() int {
	notifications := list.Notifications()
	unseen := 0

	for _, notification := range notifications {
		if notification.Seen == "" {
			unseen++
		}
	}

	return unseen
}

// Add adds an user to the list if it hasn't been added yet.
func (list *UserNotifications) Add(notificationID string) error {
	if list.Contains(notificationID) {
		return errors.New("Notification " + notificationID + " has already been added")
	}

	list.Items = append(list.Items, notificationID)
	return nil
}

// Remove removes the notification ID from the list.
func (list *UserNotifications) Remove(notificationID string) bool {
	for index, item := range list.Items {
		if item == notificationID {
			list.Items = append(list.Items[:index], list.Items[index+1:]...)
			return true
		}
	}

	return false
}

// Contains checks if the list contains the notification ID already.
func (list *UserNotifications) Contains(notificationID string) bool {
	for _, item := range list.Items {
		if item == notificationID {
			return true
		}
	}

	return false
}

// Notifications returns a slice of all the notifications.
func (list *UserNotifications) Notifications() []*Notification {
	notificationsObj := DB.GetMany("Notification", list.Items)
	notifications := []*Notification{}

	for _, obj := range notificationsObj {
		notification, ok := obj.(*Notification)

		if ok {
			notifications = append(notifications, notification)
		}
	}

	return notifications
}

// GetUserNotifications ...
func GetUserNotifications(id UserID) (*UserNotifications, error) {
	obj, err := DB.Get("UserNotifications", id)

	if err != nil {
		return nil, err
	}

	return obj.(*UserNotifications), nil
}

// StreamUserNotifications returns a stream of all user notifications.
func StreamUserNotifications() <-chan *UserNotifications {
	channel := make(chan *UserNotifications, nano.ChannelBufferSize)

	go func() {
		for obj := range DB.All("UserNotifications") {
			channel <- obj.(*UserNotifications)
		}

		close(channel)
	}()

	return channel
}

// AllUserNotifications returns a slice of all user notifications.
func AllUserNotifications() ([]*UserNotifications, error) {
	all := make([]*UserNotifications, 0, DB.Collection("UserNotifications").Count())

	for obj := range StreamUserNotifications() {
		all = append(all, obj)
	}

	return all, nil
}
