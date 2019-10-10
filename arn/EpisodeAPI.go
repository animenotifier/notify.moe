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
	_ fmt.Stringer           = (*Episode)(nil)
	_ api.Editable           = (*Episode)(nil)
	_ api.ArrayEventListener = (*Episode)(nil)
)

// Authorize returns an error if the given API POST request is not authorized.
func (episode *Episode) Authorize(ctx aero.Context, action string) error {
	user := GetUserFromContext(ctx)

	if user == nil || (user.Role != "editor" && user.Role != "admin") {
		return errors.New("Not logged in or not authorized to edit")
	}

	return nil
}

// Edit creates an edit log entry.
func (episode *Episode) Edit(ctx aero.Context, key string, value reflect.Value, newValue reflect.Value) (consumed bool, err error) {
	return edit(episode, ctx, key, value, newValue)
}

// OnAppend saves a log entry.
func (episode *Episode) OnAppend(ctx aero.Context, key string, index int, obj interface{}) {
	onAppend(episode, ctx, key, index, obj)
}

// OnRemove saves a log entry.
func (episode *Episode) OnRemove(ctx aero.Context, key string, index int, obj interface{}) {
	onRemove(episode, ctx, key, index, obj)
}

// Save saves the episode in the database.
func (episode *Episode) Save() {
	DB.Set("Episode", episode.ID, episode)
}

// Delete deletes the episode list from the database.
func (episode *Episode) Delete() error {
	DB.Delete("Episode", episode.ID)
	return nil
}
