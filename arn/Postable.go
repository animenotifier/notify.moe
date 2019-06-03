package arn

import (
	"reflect"
	"sort"
)

// Postable is a generic interface for Threads, Posts and Messages.
type Postable interface {
	Likeable

	TitleByUser(*User) string
	HTML() string
	Parent() PostParent
	Posts() []*Post
	CountPosts() int
	TypeName() string
	Creator() *User

	// Use Get prefix for these to avoid a
	// name clash with the internal fields.
	GetID() string
	GetText() string
	GetCreated() string
	GetParentID() string
}

// ToPostables converts a slice of specific types to a slice of generic postables.
func ToPostables(sliceOfPosts interface{}) []Postable {
	var postables []Postable

	v := reflect.ValueOf(sliceOfPosts)

	for i := 0; i < v.Len(); i++ {
		postable := v.Index(i).Interface().(Postable)
		postables = append(postables, postable)
	}

	return postables
}

// SortPostablesLatestFirst ...
func SortPostablesLatestFirst(posts []Postable) {
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].GetCreated() > posts[j].GetCreated()
	})
}
