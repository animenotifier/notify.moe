package arn

import (
	"github.com/aerogo/aero"
	"github.com/aerogo/api"
	jsoniter "github.com/json-iterator/go"
)

// Force interface implementations
var (
	_ api.Newable = (*Analytics)(nil)
)

// Authorize returns an error if the given API POST request is not authorized.
func (analytics *Analytics) Authorize(ctx aero.Context, action string) error {
	return AuthorizeIfLoggedIn(ctx)
}

// Create creates a new analytics object.
func (analytics *Analytics) Create(ctx aero.Context) error {
	body, err := ctx.Request().Body().Bytes()

	if err != nil {
		return err
	}

	err = jsoniter.Unmarshal(body, analytics)

	if err != nil {
		return err
	}

	analytics.UserID = GetUserFromContext(ctx).ID

	return nil
}

// Save saves the analytics in the database.
func (analytics *Analytics) Save() {
	DB.Set("Analytics", analytics.UserID, analytics)
}
