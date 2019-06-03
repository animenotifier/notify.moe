package arn

import (
	"github.com/aerogo/aero"
	"github.com/aerogo/api"
)

// Force interface implementations
var (
	_ api.Editable = (*AnimeList)(nil)
	_ IDCollection = (*AnimeList)(nil)
)

// Actions
func init() {
	API.RegisterActions("AnimeList", []*api.Action{
		// Add follow
		AddAction(),

		// Remove follow
		RemoveAction(),
	})
}

// Authorize returns an error if the given API request is not authorized.
func (list *AnimeList) Authorize(ctx aero.Context, action string) error {
	return AuthorizeIfLoggedInAndOwnData(ctx, "id")
}

// Save saves the anime list in the database.
func (list *AnimeList) Save() {
	DB.Set("AnimeList", list.UserID, list)
}
