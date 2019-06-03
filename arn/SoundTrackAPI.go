package arn

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/aerogo/aero"
	"github.com/aerogo/api"
)

// Force interface implementations
var (
	_ Publishable            = (*SoundTrack)(nil)
	_ Likeable               = (*SoundTrack)(nil)
	_ LikeEventReceiver      = (*SoundTrack)(nil)
	_ PostParent             = (*SoundTrack)(nil)
	_ fmt.Stringer           = (*SoundTrack)(nil)
	_ api.Newable            = (*SoundTrack)(nil)
	_ api.Editable           = (*SoundTrack)(nil)
	_ api.Deletable          = (*SoundTrack)(nil)
	_ api.ArrayEventListener = (*SoundTrack)(nil)
)

// Actions
func init() {
	API.RegisterActions("SoundTrack", []*api.Action{
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

// Create sets the data for a new track with data we received from the API request.
func (track *SoundTrack) Create(ctx aero.Context) error {
	user := GetUserFromContext(ctx)

	if user == nil {
		return errors.New("Not logged in")
	}

	track.ID = GenerateID("SoundTrack")
	track.Created = DateTimeUTC()
	track.CreatedBy = user.ID

	// Write log entry
	logEntry := NewEditLogEntry(user.ID, "create", "SoundTrack", track.ID, "", "", "")
	logEntry.Save()

	return track.Unpublish()
}

// Edit updates the external media object.
func (track *SoundTrack) Edit(ctx aero.Context, key string, value reflect.Value, newValue reflect.Value) (bool, error) {
	user := GetUserFromContext(ctx)

	// Write log entry
	logEntry := NewEditLogEntry(user.ID, "edit", "SoundTrack", track.ID, key, fmt.Sprint(value.Interface()), fmt.Sprint(newValue.Interface()))
	logEntry.Save()

	// Verify service name
	if strings.HasPrefix(key, "Media[") && strings.HasSuffix(key, ".Service") {
		newService := newValue.String()
		found := false

		for _, option := range DataLists["media-services"] {
			if option.Label == newService {
				found = true
				break
			}
		}

		if !found {
			return true, errors.New("Invalid service name")
		}

		value.SetString(newService)
		return true, nil
	}

	return false, nil
}

// OnAppend saves a log entry.
func (track *SoundTrack) OnAppend(ctx aero.Context, key string, index int, obj interface{}) {
	onAppend(track, ctx, key, index, obj)
}

// OnRemove saves a log entry.
func (track *SoundTrack) OnRemove(ctx aero.Context, key string, index int, obj interface{}) {
	onRemove(track, ctx, key, index, obj)
}

// DeleteInContext deletes the track in the given context.
func (track *SoundTrack) DeleteInContext(ctx aero.Context) error {
	user := GetUserFromContext(ctx)

	// Write log entry
	logEntry := NewEditLogEntry(user.ID, "delete", "SoundTrack", track.ID, "", fmt.Sprint(track), "")
	logEntry.Save()

	return track.Delete()
}

// Delete deletes the object from the database.
func (track *SoundTrack) Delete() error {
	if track.IsDraft {
		draftIndex := track.Creator().DraftIndex()
		draftIndex.SoundTrackID = ""
		draftIndex.Save()
	}

	for _, post := range track.Posts() {
		err := post.Delete()

		if err != nil {
			return err
		}
	}

	DB.Delete("SoundTrack", track.ID)
	return nil
}

// Authorize returns an error if the given API POST request is not authorized.
func (track *SoundTrack) Authorize(ctx aero.Context, action string) error {
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

// Save saves the soundtrack object in the database.
func (track *SoundTrack) Save() {
	DB.Set("SoundTrack", track.ID, track)
}
