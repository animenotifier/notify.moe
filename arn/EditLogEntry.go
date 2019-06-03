package arn

import (
	"reflect"
	"sort"

	"github.com/aerogo/nano"
)

// EditLogEntry is an entry in the editor log.
type EditLogEntry struct {
	ID         string `json:"id"`
	UserID     string `json:"userId"`
	Action     string `json:"action"`
	ObjectType string `json:"objectType"` // The typename of what was edited
	ObjectID   string `json:"objectId"`   // The ID of what was edited
	Key        string `json:"key"`
	OldValue   string `json:"oldValue"`
	NewValue   string `json:"newValue"`
	Created    string `json:"created"`
}

// NewEditLogEntry ...
func NewEditLogEntry(userID, action, objectType, objectID, key, oldValue, newValue string) *EditLogEntry {
	return &EditLogEntry{
		ID:         GenerateID("EditLogEntry"),
		UserID:     userID,
		Action:     action,
		ObjectType: objectType,
		ObjectID:   objectID,
		Key:        key,
		OldValue:   oldValue,
		NewValue:   newValue,
		Created:    DateTimeUTC(),
	}
}

// User returns the user the log entry belongs to.
func (entry *EditLogEntry) User() *User {
	user, _ := GetUser(entry.UserID)
	return user
}

// Object returns the object the log entry refers to.
func (entry *EditLogEntry) Object() interface{} {
	obj, _ := DB.Get(entry.ObjectType, entry.ObjectID)
	return obj
}

// EditorScore returns the editing score for this log entry.
func (entry *EditLogEntry) EditorScore() int {
	switch entry.Action {
	case "create":
		obj, err := DB.Get(entry.ObjectType, entry.ObjectID)

		if err != nil {
			return 0
		}

		v := reflect.Indirect(reflect.ValueOf(obj))
		isDraft := v.FieldByName("IsDraft")

		if isDraft.Kind() == reflect.Bool && isDraft.Bool() {
			// No score for drafts
			return 0
		}

		return 4

	case "edit":
		score := 4

		// Bonus score for editing anime
		if entry.ObjectType == "Anime" {
			score++

			// Bonus score for editing anime synopsis
			if entry.Key == "Summary" || entry.Key == "Synopsis" {
				score++
			}
		}

		return score

	case "delete", "arrayRemove":
		return 3

	case "arrayAppend":
		return 0
	}

	return 0
}

// ActionHumanReadable returns the human readable version of the action.
func (entry *EditLogEntry) ActionHumanReadable() string {
	switch entry.Action {
	case "create":
		return "Created"

	case "edit":
		return "Edited"

	case "delete":
		return "Deleted"

	case "arrayAppend":
		return "Added an element"

	case "arrayRemove":
		return "Removed an element"

	default:
		return entry.Action
	}
}

// StreamEditLogEntries returns a stream of all log entries.
func StreamEditLogEntries() <-chan *EditLogEntry {
	channel := make(chan *EditLogEntry, nano.ChannelBufferSize)

	go func() {
		for obj := range DB.All("EditLogEntry") {
			channel <- obj.(*EditLogEntry)
		}

		close(channel)
	}()

	return channel
}

// AllEditLogEntries returns a slice of all log entries.
func AllEditLogEntries() []*EditLogEntry {
	all := make([]*EditLogEntry, 0, DB.Collection("EditLogEntry").Count())

	stream := StreamEditLogEntries()

	for obj := range stream {
		all = append(all, obj)
	}

	return all
}

// FilterEditLogEntries filters all log entries by a custom function.
func FilterEditLogEntries(filter func(*EditLogEntry) bool) []*EditLogEntry {
	var filtered []*EditLogEntry

	channel := DB.All("EditLogEntry")

	for obj := range channel {
		realObject := obj.(*EditLogEntry)

		if filter(realObject) {
			filtered = append(filtered, realObject)
		}
	}

	return filtered
}

// SortEditLogEntriesLatestFirst puts the latest entries on top.
func SortEditLogEntriesLatestFirst(entries []*EditLogEntry) {
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Created > entries[j].Created
	})
}
