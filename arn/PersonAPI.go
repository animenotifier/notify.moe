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
	_ Likeable      = (*Person)(nil)
	_ Publishable   = (*Person)(nil)
	_ PostParent    = (*Person)(nil)
	_ fmt.Stringer  = (*Person)(nil)
	_ api.Newable   = (*Person)(nil)
	_ api.Editable  = (*Person)(nil)
	_ api.Deletable = (*Person)(nil)
)

// Actions
func init() {
	API.RegisterActions("Person", []*api.Action{
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

// Authorize returns an error if the given API request is not authorized.
func (person *Person) Authorize(ctx aero.Context, action string) error {
	user := GetUserFromContext(ctx)

	if user == nil {
		return errors.New("Not logged in")
	}

	return nil
}

// Create sets the data for a new person with data we received from the API request.
func (person *Person) Create(ctx aero.Context) error {
	user := GetUserFromContext(ctx)

	if user == nil {
		return errors.New("Not logged in")
	}

	person.ID = GenerateID("Person")
	person.Created = DateTimeUTC()
	person.CreatedBy = user.ID

	// Write log entry
	logEntry := NewEditLogEntry(user.ID, "create", "Person", person.ID, "", "", "")
	logEntry.Save()

	return person.Unpublish()
}

// Edit creates an edit log entry.
func (person *Person) Edit(ctx aero.Context, key string, value reflect.Value, newValue reflect.Value) (consumed bool, err error) {
	return edit(person, ctx, key, value, newValue)
}

// OnAppend saves a log entry.
func (person *Person) OnAppend(ctx aero.Context, key string, index int, obj interface{}) {
	onAppend(person, ctx, key, index, obj)
}

// OnRemove saves a log entry.
func (person *Person) OnRemove(ctx aero.Context, key string, index int, obj interface{}) {
	onRemove(person, ctx, key, index, obj)
}

// DeleteInContext deletes the person in the given context.
func (person *Person) DeleteInContext(ctx aero.Context) error {
	user := GetUserFromContext(ctx)

	// Write log entry
	logEntry := NewEditLogEntry(user.ID, "delete", "Person", person.ID, "", fmt.Sprint(person), "")
	logEntry.Save()

	return person.Delete()
}

// Delete deletes the object from the database.
func (person *Person) Delete() error {
	if person.IsDraft {
		draftIndex := person.Creator().DraftIndex()
		draftIndex.CharacterID = ""
		draftIndex.Save()
	}

	// Delete image files
	person.DeleteImages()

	// Delete person
	DB.Delete("Person", person.ID)
	return nil
}

// Save saves the person in the database.
func (person *Person) Save() {
	DB.Set("Person", person.ID, person)
}
