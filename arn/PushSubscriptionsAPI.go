package arn

import (
	"errors"

	"github.com/aerogo/aero"
	"github.com/aerogo/api"
	jsoniter "github.com/json-iterator/go"
)

// Force interface implementations
var (
	_ api.Editable = (*PushSubscriptions)(nil)
	_ api.Filter   = (*PushSubscriptions)(nil)
)

// Actions
func init() {
	API.RegisterActions("PushSubscriptions", []*api.Action{
		// Add subscription
		{
			Name:  "add",
			Route: "/add",
			Run: func(obj interface{}, ctx aero.Context) error {
				subscriptions := obj.(*PushSubscriptions)

				// Parse body
				body, err := ctx.Request().Body().Bytes()

				if err != nil {
					return err
				}

				var subscription *PushSubscription
				err = jsoniter.Unmarshal(body, &subscription)

				if err != nil {
					return err
				}

				// Add subscription
				err = subscriptions.Add(subscription)

				if err != nil {
					return err
				}

				subscriptions.Save()

				return nil
			},
		},

		// Remove subscription
		{
			Name:  "remove",
			Route: "/remove",
			Run: func(obj interface{}, ctx aero.Context) error {
				subscriptions := obj.(*PushSubscriptions)

				// Parse body
				body, err := ctx.Request().Body().Bytes()

				if err != nil {
					return err
				}

				var subscription *PushSubscription
				err = jsoniter.Unmarshal(body, &subscription)

				if err != nil {
					return err
				}

				// Remove subscription
				if !subscriptions.Remove(subscription.ID()) {
					return errors.New("PushSubscription does not exist")
				}

				subscriptions.Save()

				return nil
			},
		},
	})
}

// Filter removes privacy critical fields from the settings object.
func (list *PushSubscriptions) Filter() {
	for _, item := range list.Items {
		item.P256DH = ""
		item.Auth = ""
		item.Endpoint = ""
	}
}

// ShouldFilter tells whether data needs to be filtered in the given context.
func (list *PushSubscriptions) ShouldFilter(ctx aero.Context) bool {
	ctxUser := GetUserFromContext(ctx)

	if ctxUser != nil && ctxUser.Role == "admin" {
		return false
	}

	return true
}

// Authorize returns an error if the given API request is not authorized.
func (list *PushSubscriptions) Authorize(ctx aero.Context, action string) error {
	return AuthorizeIfLoggedInAndOwnData(ctx, "id")
}

// Save saves the push subscriptions in the database.
func (list *PushSubscriptions) Save() {
	DB.Set("PushSubscriptions", list.UserID, list)
}
