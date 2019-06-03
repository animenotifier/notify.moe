package arn

// import (
// 	"errors"
// 	"reflect"

// 	"github.com/aerogo/aero"
// 	"github.com/aerogo/api"
// )

// // HasEditing implements basic API functionality for editing the fields in the struct.
// type hasEditing struct {
// 	Loggable
// }

// // Force interface implementations
// var (
// 	_ api.Editable           = (*HasEditing)(nil)
// 	_ api.ArrayEventListener = (*HasEditing)(nil)
// )

// // Authorize returns an error if the given API POST request is not authorized.
// func (editable *hasEditing) Authorize(ctx aero.Context, action string) error {
// 	user := GetUserFromContext(ctx)

// 	if user == nil || (user.Role != "editor" && user.Role != "admin") {
// 		return errors.New("Not logged in or not authorized to edit")
// 	}

// 	return nil
// }

// // Edit creates an edit log entry.
// func (editable *hasEditing) Edit(ctx aero.Context, key string, value reflect.Value, newValue reflect.Value) (consumed bool, err error) {
// 	return edit(editable.Self(), ctx, key, value, newValue)
// }

// // OnAppend saves a log entry.
// func (editable *hasEditing) OnAppend(ctx aero.Context, key string, index int, obj interface{}) {
// 	onAppend(editable.Self(), ctx, key, index, obj)
// }

// // OnRemove saves a log entry.
// func (editable *hasEditing) OnRemove(ctx aero.Context, key string, index int, obj interface{}) {
// 	onRemove(editable.Self(), ctx, key, index, obj)
// }

// // Save saves the character in the database.
// func (editable *hasEditing) Save() {
// 	DB.Set(editable.TypeName(), editable.GetID(), editable.Self())
// }

// // Delete deletes the character list from the database.
// func (editable *hasEditing) Delete() error {
// 	DB.Delete(editable.TypeName(), editable.GetID())
// 	return nil
// }
