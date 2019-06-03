package arn

import (
	"github.com/aerogo/aero"
	"github.com/aerogo/api"
)

// Force interface implementations
var (
	_ IDCollection = (*UserFollows)(nil)
	_ api.Editable = (*UserFollows)(nil)
)

// Actions
func init() {
	API.RegisterActions("UserFollows", []*api.Action{
		// Add follow
		AddAction(),

		// Remove follow
		RemoveAction(),
	})
}

// Authorize returns an error if the given API request is not authorized.
func (list *UserFollows) Authorize(ctx aero.Context, action string) error {
	return AuthorizeIfLoggedInAndOwnData(ctx, "id")
}

// Save saves the follow list in the database.
func (list *UserFollows) Save() {
	DB.Set("UserFollows", list.UserID, list)
}
