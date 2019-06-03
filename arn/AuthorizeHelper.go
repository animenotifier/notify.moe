package arn

import (
	"errors"

	"github.com/aerogo/aero"
)

// AuthorizeIfLoggedInAndOwnData authorizes the given request if a user is logged in
// and the user ID matches the ID in the request.
func AuthorizeIfLoggedInAndOwnData(ctx aero.Context, userIDParameterName string) error {
	err := AuthorizeIfLoggedIn(ctx)

	if err != nil {
		return err
	}

	userID := ctx.Session().Get("userId").(string)

	if userID != ctx.Get(userIDParameterName) {
		return errors.New("Can not modify data from other users")
	}

	return nil
}

// AuthorizeIfLoggedIn authorizes the given request if a user is logged in.
func AuthorizeIfLoggedIn(ctx aero.Context) error {
	if !ctx.HasSession() {
		return errors.New("Neither logged in nor in session")
	}

	userID, ok := ctx.Session().Get("userId").(string)

	if !ok || userID == "" {
		return errors.New("Not logged in")
	}

	return nil
}
