package arn

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/aerogo/aero"
	"github.com/aerogo/api"
)

// Force interface implementations
var (
	_ Likeable          = (*Quote)(nil)
	_ LikeEventReceiver = (*Quote)(nil)
	_ Publishable       = (*Quote)(nil)
	_ PostParent        = (*Quote)(nil)
	_ fmt.Stringer      = (*Quote)(nil)
	_ api.Newable       = (*Quote)(nil)
	_ api.Editable      = (*Quote)(nil)
	_ api.Deletable     = (*Quote)(nil)
)

// Actions
func init() {
	API.RegisterActions("Quote", []*api.Action{
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

// Create sets the data for a new quote with data we received from the API request.
func (quote *Quote) Create(ctx aero.Context) error {
	user := GetUserFromContext(ctx)

	if user == nil {
		return errors.New("Not logged in")
	}

	quote.ID = GenerateID("Quote")
	quote.Created = DateTimeUTC()
	quote.CreatedBy = user.ID
	quote.EpisodeNumber = -1
	quote.Time = -1

	// Write log entry
	logEntry := NewEditLogEntry(user.ID, "create", "Quote", quote.ID, "", "", "")
	logEntry.Save()

	return quote.Unpublish()
}

// Edit saves a log entry for the edit.
func (quote *Quote) Edit(ctx aero.Context, key string, value reflect.Value, newValue reflect.Value) (bool, error) {
	user := GetUserFromContext(ctx)

	// Write log entry
	logEntry := NewEditLogEntry(user.ID, "edit", "Quote", quote.ID, key, fmt.Sprint(value.Interface()), fmt.Sprint(newValue.Interface()))
	logEntry.Save()

	return false, nil
}

// Save saves the quote in the database.
func (quote *Quote) Save() {
	DB.Set("Quote", quote.ID, quote)
}

// DeleteInContext deletes the quote in the given context.
func (quote *Quote) DeleteInContext(ctx aero.Context) error {
	user := GetUserFromContext(ctx)

	// Write log entry
	logEntry := NewEditLogEntry(user.ID, "delete", "Quote", quote.ID, "", fmt.Sprint(quote), "")
	logEntry.Save()

	return quote.Delete()
}

// Delete deletes the object from the database.
func (quote *Quote) Delete() error {
	if quote.IsDraft {
		draftIndex := quote.Creator().DraftIndex()
		draftIndex.QuoteID = ""
		draftIndex.Save()
	}

	// Remove posts
	for _, post := range quote.Posts() {
		err := post.Delete()

		if err != nil {
			return err
		}
	}

	// Remove main quote reference
	character := quote.Character()

	if character.MainQuoteID == quote.ID {
		character.MainQuoteID = ""
		character.Save()
	}

	DB.Delete("Quote", quote.ID)
	return nil
}

// Authorize returns an error if the given API request is not authorized.
func (quote *Quote) Authorize(ctx aero.Context, action string) error {
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
