package arn

import (
	"errors"

	"github.com/aerogo/aero"
	"github.com/aerogo/api"
)

// IDCollection ...
type IDCollection interface {
	Add(id string) error
	Remove(id string) bool
	Save()
}

// AddAction returns an API action that adds a new item to the IDCollection.
func AddAction() *api.Action {
	return &api.Action{
		Name:  "add",
		Route: "/add/:item-id",
		Run: func(obj interface{}, ctx aero.Context) error {
			list := obj.(IDCollection)
			itemID := ctx.Get("item-id")
			err := list.Add(itemID)

			if err != nil {
				return err
			}

			list.Save()
			return nil
		},
	}
}

// RemoveAction returns an API action that removes an item from the IDCollection.
func RemoveAction() *api.Action {
	return &api.Action{
		Name:  "remove",
		Route: "/remove/:item-id",
		Run: func(obj interface{}, ctx aero.Context) error {
			list := obj.(IDCollection)
			itemID := ctx.Get("item-id")

			if !list.Remove(itemID) {
				return errors.New("This item does not exist in the list")
			}

			list.Save()
			return nil
		},
	}
}
