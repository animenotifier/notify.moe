package arn

import (
	"fmt"
	"reflect"

	"github.com/aerogo/aero"
)

// Loggable applies to any type that has a TypeName function.
type Loggable interface {
	GetID() string
	TypeName() string
	Self() Loggable
}

// edit creates an edit log entry.
func edit(loggable Loggable, ctx aero.Context, key string, value reflect.Value, newValue reflect.Value) (consumed bool, err error) {
	user := GetUserFromContext(ctx)

	// Write log entry
	logEntry := NewEditLogEntry(user.ID, "edit", loggable.TypeName(), loggable.GetID(), key, fmt.Sprint(value.Interface()), fmt.Sprint(newValue.Interface()))
	logEntry.Save()

	return false, nil
}

// onAppend saves a log entry.
func onAppend(loggable Loggable, ctx aero.Context, key string, index int, obj interface{}) {
	user := GetUserFromContext(ctx)
	logEntry := NewEditLogEntry(user.ID, "arrayAppend", loggable.TypeName(), loggable.GetID(), fmt.Sprintf("%s[%d]", key, index), "", fmt.Sprint(obj))
	logEntry.Save()
}

// onRemove saves a log entry.
func onRemove(loggable Loggable, ctx aero.Context, key string, index int, obj interface{}) {
	user := GetUserFromContext(ctx)
	logEntry := NewEditLogEntry(user.ID, "arrayRemove", loggable.TypeName(), loggable.GetID(), fmt.Sprintf("%s[%d]", key, index), fmt.Sprint(obj), "")
	logEntry.Save()
}
