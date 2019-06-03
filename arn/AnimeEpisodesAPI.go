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
	_ fmt.Stringer           = (*AnimeEpisodes)(nil)
	_ api.Editable           = (*AnimeEpisodes)(nil)
	_ api.ArrayEventListener = (*AnimeEpisodes)(nil)
)

// Authorize returns an error if the given API POST request is not authorized.
func (episodes *AnimeEpisodes) Authorize(ctx aero.Context, action string) error {
	user := GetUserFromContext(ctx)

	if user == nil || (user.Role != "editor" && user.Role != "admin") {
		return errors.New("Not logged in or not authorized to edit")
	}

	return nil
}

// Edit creates an edit log entry.
func (episodes *AnimeEpisodes) Edit(ctx aero.Context, key string, value reflect.Value, newValue reflect.Value) (consumed bool, err error) {
	return edit(episodes, ctx, key, value, newValue)
}

// OnAppend saves a log entry.
func (episodes *AnimeEpisodes) OnAppend(ctx aero.Context, key string, index int, obj interface{}) {
	onAppend(episodes, ctx, key, index, obj)
}

// OnRemove saves a log entry.
func (episodes *AnimeEpisodes) OnRemove(ctx aero.Context, key string, index int, obj interface{}) {
	onRemove(episodes, ctx, key, index, obj)
}

// Save saves the episodes in the database.
func (episodes *AnimeEpisodes) Save() {
	DB.Set("AnimeEpisodes", episodes.AnimeID, episodes)
}

// Delete deletes the episode list from the database.
func (episodes *AnimeEpisodes) Delete() error {
	DB.Delete("AnimeEpisodes", episodes.AnimeID)
	return nil
}
