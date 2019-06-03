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
	_ fmt.Stringer           = (*AnimeRelations)(nil)
	_ api.Editable           = (*AnimeRelations)(nil)
	_ api.ArrayEventListener = (*AnimeRelations)(nil)
)

// Authorize returns an error if the given API POST request is not authorized.
func (relations *AnimeRelations) Authorize(ctx aero.Context, action string) error {
	user := GetUserFromContext(ctx)

	if user == nil || (user.Role != "editor" && user.Role != "admin") {
		return errors.New("Not logged in or not authorized to edit")
	}

	return nil
}

// Edit creates an edit log entry.
func (relations *AnimeRelations) Edit(ctx aero.Context, key string, value reflect.Value, newValue reflect.Value) (consumed bool, err error) {
	return edit(relations, ctx, key, value, newValue)
}

// OnAppend saves a log entry.
func (relations *AnimeRelations) OnAppend(ctx aero.Context, key string, index int, obj interface{}) {
	onAppend(relations, ctx, key, index, obj)
}

// OnRemove saves a log entry.
func (relations *AnimeRelations) OnRemove(ctx aero.Context, key string, index int, obj interface{}) {
	onRemove(relations, ctx, key, index, obj)
}

// Save saves the anime relations object in the database.
func (relations *AnimeRelations) Save() {
	DB.Set("AnimeRelations", relations.AnimeID, relations)
}

// Delete deletes the relation list from the database.
func (relations *AnimeRelations) Delete() error {
	DB.Delete("AnimeRelations", relations.AnimeID)
	return nil
}
