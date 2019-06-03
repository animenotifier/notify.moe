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
	_ Likeable      = (*Character)(nil)
	_ Publishable   = (*Character)(nil)
	_ PostParent    = (*Character)(nil)
	_ fmt.Stringer  = (*Character)(nil)
	_ api.Newable   = (*Character)(nil)
	_ api.Editable  = (*Character)(nil)
	_ api.Deletable = (*Character)(nil)
)

// Actions
func init() {
	API.RegisterActions("Character", []*api.Action{
		// Publish
		PublishAction(),

		// Unpublish
		UnpublishAction(),

		// Like character
		LikeAction(),

		// Unlike character
		UnlikeAction(),
	})
}

// Create sets the data for a new character with data we received from the API request.
func (character *Character) Create(ctx aero.Context) error {
	user := GetUserFromContext(ctx)

	if user == nil {
		return errors.New("Not logged in")
	}

	character.ID = GenerateID("Character")
	character.Created = DateTimeUTC()
	character.CreatedBy = user.ID

	// Write log entry
	logEntry := NewEditLogEntry(user.ID, "create", "Character", character.ID, "", "", "")
	logEntry.Save()

	return character.Unpublish()
}

// Authorize returns an error if the given API request is not authorized.
func (character *Character) Authorize(ctx aero.Context, action string) error {
	user := GetUserFromContext(ctx)

	if user == nil {
		return errors.New("Not logged in")
	}

	// Allow custom actions (like, unlike) for normal users
	if action == "like" || action == "unlike" {
		return nil
	}

	if user.Role != "editor" && user.Role != "admin" {
		return errors.New("Insufficient permissions")
	}

	return nil
}

// Edit creates an edit log entry.
func (character *Character) Edit(ctx aero.Context, key string, value reflect.Value, newValue reflect.Value) (consumed bool, err error) {
	return edit(character, ctx, key, value, newValue)
}

// OnAppend saves a log entry.
func (character *Character) OnAppend(ctx aero.Context, key string, index int, obj interface{}) {
	onAppend(character, ctx, key, index, obj)
}

// OnRemove saves a log entry.
func (character *Character) OnRemove(ctx aero.Context, key string, index int, obj interface{}) {
	onRemove(character, ctx, key, index, obj)
}

// DeleteInContext deletes the character in the given context.
func (character *Character) DeleteInContext(ctx aero.Context) error {
	user := GetUserFromContext(ctx)

	// Write log entry
	logEntry := NewEditLogEntry(user.ID, "delete", "Character", character.ID, "", fmt.Sprint(character), "")
	logEntry.Save()

	return character.Delete()
}

// Delete deletes the object from the database.
func (character *Character) Delete() error {
	if character.IsDraft {
		draftIndex := character.Creator().DraftIndex()
		draftIndex.CharacterID = ""
		draftIndex.Save()
	}

	// Delete from anime characters
	for list := range StreamAnimeCharacters() {
		list.Lock()

		for index, animeCharacter := range list.Items {
			if animeCharacter.CharacterID == character.ID {
				list.Items = append(list.Items[:index], list.Items[index+1:]...)
				list.Save()
				break
			}
		}

		list.Unlock()
	}

	// Delete from quotes
	for quote := range StreamQuotes() {
		if quote.CharacterID == character.ID {
			err := quote.Delete()

			if err != nil {
				return err
			}
		}
	}

	// Delete image files
	character.DeleteImages()

	// Delete character
	DB.Delete("Character", character.ID)
	return nil
}

// Save saves the character in the database.
func (character *Character) Save() {
	DB.Set("Character", character.ID, character)
}
