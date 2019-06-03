package arn

import (
	"errors"
	"fmt"
	"os"
	"path"
	"sync"

	"github.com/aerogo/nano"
	"github.com/akyoto/color"
)

// Group represents a group of users.
type Group struct {
	Name        string         `json:"name" editable:"true"`
	Tagline     string         `json:"tagline" editable:"true"`
	Image       GroupImage     `json:"image"`
	Description string         `json:"description" editable:"true" type:"textarea"`
	Rules       string         `json:"rules" editable:"true" type:"textarea"`
	Restricted  bool           `json:"restricted" editable:"true" tooltip:"Restricted groups can only be joined with the founder's permission."`
	Tags        []string       `json:"tags" editable:"true"`
	Members     []*GroupMember `json:"members"`
	Neighbors   []string       `json:"neighbors"`
	// Applications []UserApplication `json:"applications"`

	// Mixins
	hasID
	hasPosts
	hasCreator
	hasEditor
	hasDraft

	// Mutex
	membersMutex sync.Mutex
}

// Link returns the URI to the group page.
func (group *Group) Link() string {
	return "/group/" + group.ID
}

// TitleByUser returns the preferred title for the given user.
func (group *Group) TitleByUser(user *User) string {
	if group.Name == "" {
		return "untitled"
	}

	return group.Name
}

// String is the default text representation of the group.
func (group *Group) String() string {
	return group.TitleByUser(nil)
}

// FindMember returns the group member by user ID, if available.
func (group *Group) FindMember(userID UserID) *GroupMember {
	group.membersMutex.Lock()
	defer group.membersMutex.Unlock()

	for _, member := range group.Members {
		if member.UserID == userID {
			return member
		}
	}

	return nil
}

// HasMember returns true if the user is a member of the group.
func (group *Group) HasMember(userID UserID) bool {
	return group.FindMember(userID) != nil
}

// Users returns a slice of all users in the group.
func (group *Group) Users() []*User {
	group.membersMutex.Lock()
	defer group.membersMutex.Unlock()
	users := make([]*User, len(group.Members))

	for index, member := range group.Members {
		users[index] = member.User()
	}

	return users
}

// TypeName returns the type name.
func (group *Group) TypeName() string {
	return "Group"
}

// Self returns the object itself.
func (group *Group) Self() Loggable {
	return group
}

// Publish ...
func (group *Group) Publish() error {
	if len(group.Name) < 2 {
		return errors.New("Name too short: Should be at least 2 characters")
	}

	if len(group.Name) > 35 {
		return errors.New("Name too long: Should not be more than 35 characters")
	}

	if len(group.Tagline) < 4 {
		return errors.New("Tagline too short: Should be at least 4 characters")
	}

	if len(group.Tagline) > 60 {
		return errors.New("Tagline too long: Should not be more than 60 characters")
	}

	if len(group.Description) < 10 {
		return errors.New("Your group needs a description (at least 10 characters)")
	}

	if len(group.Tags) < 1 {
		return errors.New("At least one tag is required")
	}

	if !group.HasImage() {
		return errors.New("Group image required")
	}

	return publish(group)
}

// Unpublish ...
func (group *Group) Unpublish() error {
	return unpublish(group)
}

// Join makes the given user join the group.
func (group *Group) Join(user *User) error {
	// Check if the user is already a member
	member := group.FindMember(user.ID)

	if member != nil {
		return errors.New("Already a member of this group")
	}

	// Add user to the members list
	group.membersMutex.Lock()

	group.Members = append(group.Members, &GroupMember{
		UserID: user.ID,
		Joined: DateTimeUTC(),
	})

	group.membersMutex.Unlock()

	// Trigger notifications
	group.OnJoin(user)
	return nil
}

// Leave makes the given user leave the group.
func (group *Group) Leave(user *User) error {
	group.membersMutex.Lock()
	defer group.membersMutex.Unlock()

	for index, member := range group.Members {
		if member.UserID == user.ID {
			if member.UserID == group.CreatedBy {
				return errors.New("The founder can not leave the group, please contact a staff member")
			}

			group.Members = append(group.Members[:index], group.Members[index+1:]...)
			return nil
		}
	}

	return nil
}

// OnJoin sends notifications to the creator.
func (group *Group) OnJoin(user *User) {
	go func() {
		group.Creator().SendNotification(&PushNotification{
			Title:   fmt.Sprintf(`%s joined your group!`, user.Nick),
			Message: fmt.Sprintf(`%s has joined your group "%s"`, user.Nick, group.Name),
			Icon:    "https:" + user.AvatarLink("large"),
			Link:    "https://notify.moe" + group.Link() + "/members",
			Type:    NotificationTypeGroupJoin,
		})
	}()
}

// SendNotification sends a notification to all group members except for the excluded user ID.
func (group *Group) SendNotification(notification *PushNotification, excludeUserID UserID) {
	for _, user := range group.Users() {
		if user.ID == excludeUserID {
			continue
		}

		user.SendNotification(notification)
	}
}

// AverageColor returns the average color of the image.
func (group *Group) AverageColor() string {
	color := group.Image.AverageColor

	if color.Hue == 0 && color.Saturation == 0 && color.Lightness == 0 {
		return ""
	}

	return color.String()
}

// ImageLink returns a link to the group image.
func (group *Group) ImageLink(size string) string {
	if !group.HasImage() {
		return fmt.Sprintf("//%s/images/elements/no-group-image.svg", MediaHost)
	}

	extension := ".jpg"

	if size == "original" {
		extension = group.Image.Extension
	}

	return fmt.Sprintf("//%s/images/groups/%s/%s%s?%v", MediaHost, size, group.ID, extension, group.Image.LastModified)
}

// DeleteImages deletes all images for the group.
func (group *Group) DeleteImages() {
	if group.Image.Extension == "" {
		return
	}

	// Original
	err := os.Remove(path.Join(Root, "images/groups/original/", group.ID+group.Image.Extension))

	if err != nil {
		// Don't return the error.
		// It's too late to stop the process at this point.
		// Instead, log the error.
		color.Red(err.Error())
	}

	// Small
	os.Remove(path.Join(Root, "images/groups/small/", group.ID+".jpg"))
	os.Remove(path.Join(Root, "images/groups/small/", group.ID+"@2.jpg"))
	os.Remove(path.Join(Root, "images/groups/small/", group.ID+".webp"))
	os.Remove(path.Join(Root, "images/groups/small/", group.ID+"@2.webp"))

	// Large
	os.Remove(path.Join(Root, "images/groups/large/", group.ID+".jpg"))
	os.Remove(path.Join(Root, "images/groups/large/", group.ID+"@2.jpg"))
	os.Remove(path.Join(Root, "images/groups/large/", group.ID+".webp"))
	os.Remove(path.Join(Root, "images/groups/large/", group.ID+"@2.webp"))
}

// GetGroup ...
func GetGroup(id string) (*Group, error) {
	obj, err := DB.Get("Group", id)

	if err != nil {
		return nil, err
	}

	return obj.(*Group), nil
}

// StreamGroups returns a stream of all groups.
func StreamGroups() <-chan *Group {
	channel := make(chan *Group, nano.ChannelBufferSize)

	go func() {
		for obj := range DB.All("Group") {
			channel <- obj.(*Group)
		}

		close(channel)
	}()

	return channel
}

// AllGroups returns a slice of all groups.
func AllGroups() []*Group {
	all := make([]*Group, 0, DB.Collection("Group").Count())
	stream := StreamGroups()

	for obj := range stream {
		all = append(all, obj)
	}

	return all
}

// FilterGroups filters all groups by a custom function.
func FilterGroups(filter func(*Group) bool) []*Group {
	var filtered []*Group

	for obj := range DB.All("Group") {
		realObject := obj.(*Group)

		if filter(realObject) {
			filtered = append(filtered, realObject)
		}
	}

	return filtered
}
