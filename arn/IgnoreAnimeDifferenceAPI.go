package arn

import (
	"errors"
	"strconv"

	"github.com/aerogo/aero"
	"github.com/aerogo/api"
)

// Force interface implementations
var (
	_ api.Newable = (*IgnoreAnimeDifference)(nil)
)

// Authorize returns an error if the given API POST request is not authorized.
func (ignore *IgnoreAnimeDifference) Authorize(ctx aero.Context, action string) error {
	user := GetUserFromContext(ctx)

	if user == nil {
		return errors.New("Not logged in")
	}

	if user.Role != "editor" && user.Role != "admin" {
		return errors.New("Not authorized")
	}

	return nil
}

// Create constructs the values for this new object with the data we received from the API request.
func (ignore *IgnoreAnimeDifference) Create(ctx aero.Context) error {
	data, err := ctx.Request().Body().JSONObject()

	if err != nil {
		return err
	}

	user := GetUserFromContext(ctx)

	if user == nil {
		return errors.New("Not logged in")
	}

	hash, err := strconv.ParseUint(data["hash"].(string), 10, 64)

	if err != nil {
		return errors.New("Invalid hash: Not a number")
	}

	ignore.ID = data["id"].(string)
	ignore.ValueHash = hash
	ignore.Created = DateTimeUTC()
	ignore.CreatedBy = user.ID

	if ignore.ID == "" {
		return errors.New("Invalid ID")
	}

	return nil
}

// Save saves the object in the database.
func (ignore *IgnoreAnimeDifference) Save() {
	DB.Set("IgnoreAnimeDifference", ignore.ID, ignore)
}
