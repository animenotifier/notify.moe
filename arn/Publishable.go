package arn

import (
	"errors"
	"reflect"

	"github.com/aerogo/aero"
	"github.com/aerogo/api"
)

// Publishable ...
type Publishable interface {
	Publish() error
	Unpublish() error
	Save()
	GetID() string
	GetCreatedBy() string
	GetIsDraft() bool
	SetIsDraft(bool)
}

// PublishAction returns an API action that publishes the object.
func PublishAction() *api.Action {
	return &api.Action{
		Name:  "publish",
		Route: "/publish",
		Run: func(obj interface{}, ctx aero.Context) error {
			draft := obj.(Publishable)
			err := draft.Publish()

			if err != nil {
				return err
			}

			draft.Save()
			return nil
		},
	}
}

// UnpublishAction returns an API action that unpublishes the object.
func UnpublishAction() *api.Action {
	return &api.Action{
		Name:  "unpublish",
		Route: "/unpublish",
		Run: func(obj interface{}, ctx aero.Context) error {
			draft := obj.(Publishable)
			err := draft.Unpublish()

			if err != nil {
				return err
			}

			draft.Save()
			return nil
		},
	}
}

// publish is the generic publish function.
func publish(draft Publishable) error {
	// No draft
	if !draft.GetIsDraft() {
		return errors.New("Not a draft")
	}

	// Get object type
	typ := reflect.TypeOf(draft)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	// Get draft index
	draftIndex, err := GetDraftIndex(draft.GetCreatedBy())

	if err != nil {
		return err
	}

	currentDraftID, _ := draftIndex.GetID(typ.Name())

	if currentDraftID != draft.GetID() {
		return errors.New(typ.Name() + " draft doesn't exist in the user draft index")
	}

	// Publish the object
	draft.SetIsDraft(false)
	err = draftIndex.SetID(typ.Name(), "")

	if err != nil {
		return err
	}

	draftIndex.Save()

	return nil
}

// unpublish turns the object back into a draft.
func unpublish(draft Publishable) error {
	// Get object type
	typ := reflect.TypeOf(draft)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	// Get draft index
	draftIndex, err := GetDraftIndex(draft.GetCreatedBy())

	if err != nil {
		return err
	}

	draftIndexID, _ := draftIndex.GetID(typ.Name())

	if draftIndexID != "" {
		return errors.New("You still have an unfinished draft")
	}

	draft.SetIsDraft(true)
	err = draftIndex.SetID(typ.Name(), draft.GetID())

	if err != nil {
		return err
	}

	draftIndex.Save()
	return nil
}
