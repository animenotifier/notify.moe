package arn

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/aerogo/aero"
	"github.com/aerogo/aero/event"
	"github.com/aerogo/api"
	"github.com/aerogo/markdown"
	"github.com/animenotifier/notify.moe/arn/autocorrect"
	"github.com/animenotifier/notify.moe/arn/limits"
)

// Force interface implementations
var (
	_ Postable          = (*Post)(nil)
	_ Likeable          = (*Post)(nil)
	_ LikeEventReceiver = (*Post)(nil)
	_ PostParent        = (*Post)(nil)
	_ fmt.Stringer      = (*Post)(nil)
	_ api.Newable       = (*Post)(nil)
	_ api.Editable      = (*Post)(nil)
	_ api.Actionable    = (*Post)(nil)
	_ api.Deletable     = (*Post)(nil)
)

// Actions
func init() {
	API.RegisterActions("Post", []*api.Action{
		// Like post
		LikeAction(),

		// Unlike post
		UnlikeAction(),
	})
}

// Authorize returns an error if the given API POST request is not authorized.
func (post *Post) Authorize(ctx aero.Context, action string) error {
	if !ctx.HasSession() {
		return errors.New("Neither logged in nor in session")
	}

	if action == "edit" {
		user := GetUserFromContext(ctx)

		if post.CreatedBy != user.ID && user.Role != "admin" {
			return errors.New("Can't edit the posts of other users")
		}
	}

	return nil
}

// Create sets the data for a new post with data we received from the API request.
func (post *Post) Create(ctx aero.Context) error {
	data, err := ctx.Request().Body().JSONObject()

	if err != nil {
		return err
	}

	user := GetUserFromContext(ctx)

	if user == nil {
		return errors.New("Not logged in")
	}

	post.ID = GenerateID("Post")
	post.Text, _ = data["text"].(string)
	post.CreatedBy = user.ID
	post.ParentID, _ = data["parentId"].(string)
	post.ParentType, _ = data["parentType"].(string)
	post.Created = DateTimeUTC()
	post.Edited = ""

	// Check parent type
	if !DB.HasType(post.ParentType) {
		return errors.New("Invalid parent type: " + post.ParentType)
	}

	// Post-process text
	post.Text = autocorrect.PostText(post.Text)

	if len(post.Text) < limits.PostMinCharacters {
		return fmt.Errorf("Text too short: Should be at least %d characters", limits.PostMinCharacters)
	}

	// Tags
	tags, _ := data["tags"].([]interface{})
	post.Tags = make([]string, len(tags))

	for i := range post.Tags {
		post.Tags[i] = tags[i].(string)
	}

	// Parent
	parent := post.Parent()

	if parent == nil {
		return errors.New(post.ParentType + " does not exist")
	}

	// Is the parent locked?
	if IsLocked(parent) {
		return errors.New(post.ParentType + " is locked")
	}

	// Don't allow posting when you're not a group member
	topMostParent := post.TopMostParent()

	if topMostParent.TypeName() == "Group" {
		group := topMostParent.(*Group)

		if !group.HasMember(user.ID) {
			return errors.New("Only group members can post in groups")
		}
	}

	// Append to posts
	parent.AddPost(post.ID)

	// Save the parent thread
	parent.Save()

	// Send notification to the author of the parent post
	go func() {
		notifyUser := parent.Creator()

		// Does the parent have a creator?
		if notifyUser == nil {
			return
		}

		// Don't notify the author himself
		if notifyUser.ID == post.CreatedBy {
			return
		}

		title := user.Nick + " replied"
		message := ""

		switch post.ParentType {
		case "Post":
			message = fmt.Sprintf("%s replied to your comment in \"%s\".", user.Nick, parent.(*Post).Parent().TitleByUser(notifyUser))
		case "User":
			title = fmt.Sprintf("%s wrote a comment on your profile.", user.Nick)
			message = post.Text
		case "Group":
			title = fmt.Sprintf(`%s wrote a new post in the group "%s".`, user.Nick, parent.TitleByUser(nil))
			message = post.Text
		default:
			message = fmt.Sprintf(`%s replied in the %s "%s".`, user.Nick, strings.ToLower(post.ParentType), parent.TitleByUser(notifyUser))
		}

		notification := &PushNotification{
			Title:   title,
			Message: message,
			Icon:    "https:" + user.AvatarLink("large"),
			Link:    post.Link(),
			Type:    NotificationTypeForumReply,
		}

		// If you're posting to a group,
		// all members except the author will receive a notification.
		if post.ParentType == "Group" {
			group := parent.(*Group)
			group.SendNotification(notification, user.ID)
			return
		}

		notifyUser.SendNotification(notification)
	}()

	// Write log entry
	logEntry := NewEditLogEntry(user.ID, "create", "Post", post.ID, "", "", "")
	logEntry.Save()

	// Create activity
	activity := NewActivityCreate("Post", post.ID, user.ID)
	activity.Save()

	// Broadcast event to all users so they can reload the activity page if needed
	for receiver := range StreamUsers() {
		activityEvent := event.New("post activity", receiver.IsFollowing(user.ID))
		receiver.BroadcastEvent(activityEvent)
	}

	return nil
}

// Edit saves a log entry for the edit.
func (post *Post) Edit(ctx aero.Context, key string, value reflect.Value, newValue reflect.Value) (bool, error) {
	consumed := false
	user := GetUserFromContext(ctx)

	// This should stay a switch statement.
	// nolint:gocritic
	switch key {
	case "ParentID":
		var newParent PostParent
		newParentID := newValue.String()
		newParent, err := GetPost(newParentID)

		if err != nil {
			newParent, err = GetThread(newParentID)

			if err != nil {
				return false, err
			}
		}

		post.SetParent(newParent)
		consumed = true

	case "Text":
		newText := newValue.String()
		newText = autocorrect.PostText(newText)

		if len(newText) < limits.PostMinCharacters {
			return false, fmt.Errorf("Text too short: Should be at least %d characters", limits.PostMinCharacters)
		}

		post.Text = newText
		consumed = true
	}

	// Write log entry
	logEntry := NewEditLogEntry(user.ID, "edit", "Post", post.ID, key, fmt.Sprint(value.Interface()), fmt.Sprint(newValue.Interface()))
	logEntry.Save()

	return consumed, nil
}

// OnAppend saves a log entry.
func (post *Post) OnAppend(ctx aero.Context, key string, index int, obj interface{}) {
	onAppend(post, ctx, key, index, obj)
}

// OnRemove saves a log entry.
func (post *Post) OnRemove(ctx aero.Context, key string, index int, obj interface{}) {
	onRemove(post, ctx, key, index, obj)
}

// AfterEdit sets the edited date on the post object.
func (post *Post) AfterEdit(ctx aero.Context) error {
	post.Edited = DateTimeUTC()
	post.html = markdown.Render(post.Text)
	return nil
}

// DeleteInContext deletes the post in the given context.
func (post *Post) DeleteInContext(ctx aero.Context) error {
	user := GetUserFromContext(ctx)

	// Write log entry
	logEntry := NewEditLogEntry(user.ID, "delete", "Post", post.ID, "", fmt.Sprint(post), "")
	logEntry.Save()

	return post.Delete()
}

// Delete deletes the post from the database.
func (post *Post) Delete() error {
	// Remove child posts first
	for _, post := range post.Posts() {
		err := post.Delete()

		if err != nil {
			return err
		}
	}

	parent := post.Parent()

	if parent == nil {
		return fmt.Errorf("Invalid %s parent ID: %s", post.ParentType, post.ParentID)
	}

	// Remove the reference of the post in the thread that contains it
	if !parent.RemovePost(post.ID) {
		return fmt.Errorf("This post does not exist in the %s", strings.ToLower(post.ParentType))
	}

	parent.Save()

	// Remove activities
	for activity := range StreamActivityCreates() {
		if activity.ObjectID == post.ID && activity.ObjectType == "Post" {
			err := activity.Delete()

			if err != nil {
				return err
			}
		}
	}

	DB.Delete("Post", post.ID)
	return nil
}

// Save saves the post object in the database.
func (post *Post) Save() {
	DB.Set("Post", post.ID, post)
}
