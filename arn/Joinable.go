package arn

import (
	"github.com/aerogo/aero"
	"github.com/aerogo/api"
)

// Joinable is an object where users can join and leave.
type Joinable interface {
	Join(*User) error
	Leave(*User) error
	Save()
}

// JoinAction returns an API action that lets the user join the object.
func JoinAction() *api.Action {
	return &api.Action{
		Name:  "join",
		Route: "/join",
		Run: func(obj interface{}, ctx aero.Context) error {
			user := GetUserFromContext(ctx)
			joinable := obj.(Joinable)
			err := joinable.Join(user)

			if err != nil {
				return err
			}

			joinable.Save()
			return nil
		},
	}
}

// LeaveAction returns an API action that unpublishes the object.
func LeaveAction() *api.Action {
	return &api.Action{
		Name:  "leave",
		Route: "/leave",
		Run: func(obj interface{}, ctx aero.Context) error {
			user := GetUserFromContext(ctx)
			joinable := obj.(Joinable)
			err := joinable.Leave(user)

			if err != nil {
				return err
			}

			joinable.Save()
			return nil
		},
	}
}
