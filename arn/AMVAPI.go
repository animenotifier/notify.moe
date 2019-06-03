package arn

import (
	"errors"
	"fmt"
	"os"
	"path"
	"reflect"

	"github.com/aerogo/aero"
	"github.com/aerogo/api"
)

// Force interface implementations
var (
	_ Publishable            = (*AMV)(nil)
	_ Likeable               = (*AMV)(nil)
	_ LikeEventReceiver      = (*AMV)(nil)
	_ PostParent             = (*AMV)(nil)
	_ fmt.Stringer           = (*AMV)(nil)
	_ api.Newable            = (*AMV)(nil)
	_ api.Editable           = (*AMV)(nil)
	_ api.Deletable          = (*AMV)(nil)
	_ api.ArrayEventListener = (*AMV)(nil)
)

// Actions
func init() {
	API.RegisterActions("AMV", []*api.Action{
		// Publish
		PublishAction(),

		// Unpublish
		UnpublishAction(),

		// Like
		LikeAction(),

		// Unlike
		UnlikeAction(),
	})
}

// Create sets the data for a new AMV with data we received from the API request.
func (amv *AMV) Create(ctx aero.Context) error {
	user := GetUserFromContext(ctx)

	if user == nil {
		return errors.New("Not logged in")
	}

	amv.ID = GenerateID("AMV")
	amv.Created = DateTimeUTC()
	amv.CreatedBy = user.ID

	// Write log entry
	logEntry := NewEditLogEntry(user.ID, "create", "AMV", amv.ID, "", "", "")
	logEntry.Save()

	return amv.Unpublish()
}

// Edit creates an edit log entry.
func (amv *AMV) Edit(ctx aero.Context, key string, value reflect.Value, newValue reflect.Value) (consumed bool, err error) {
	return edit(amv, ctx, key, value, newValue)
}

// OnAppend saves a log entry.
func (amv *AMV) OnAppend(ctx aero.Context, key string, index int, obj interface{}) {
	onAppend(amv, ctx, key, index, obj)
}

// OnRemove saves a log entry.
func (amv *AMV) OnRemove(ctx aero.Context, key string, index int, obj interface{}) {
	onRemove(amv, ctx, key, index, obj)
}

// DeleteInContext deletes the amv in the given context.
func (amv *AMV) DeleteInContext(ctx aero.Context) error {
	user := GetUserFromContext(ctx)

	// Write log entry
	logEntry := NewEditLogEntry(user.ID, "delete", "AMV", amv.ID, "", fmt.Sprint(amv), "")
	logEntry.Save()

	return amv.Delete()
}

// Delete deletes the object from the database.
func (amv *AMV) Delete() error {
	if amv.IsDraft {
		draftIndex := amv.Creator().DraftIndex()
		draftIndex.AMVID = ""
		draftIndex.Save()
	}

	// Remove posts
	for _, post := range amv.Posts() {
		err := post.Delete()

		if err != nil {
			return err
		}
	}

	// Remove file
	if amv.File != "" {
		err := os.Remove(path.Join(Root, "videos", "amvs", amv.File))

		if err != nil {
			return err
		}
	}

	DB.Delete("AMV", amv.ID)
	return nil
}

// Authorize returns an error if the given API POST request is not authorized.
func (amv *AMV) Authorize(ctx aero.Context, action string) error {
	user := GetUserFromContext(ctx)

	if user == nil {
		return errors.New("Not logged in")
	}

	if action == "delete" {
		if user.Role != "editor" && user.Role != "admin" {
			return errors.New("Insufficient permissions")
		}
	}

	return nil
}

// Save saves the amv object in the database.
func (amv *AMV) Save() {
	DB.Set("AMV", amv.ID, amv)
}
