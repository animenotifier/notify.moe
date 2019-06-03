package arn

import (
	"github.com/aerogo/api"
)

// PostParent is an interface that defines common functions for parent objects of posts.
type PostParent interface {
	Linkable
	api.Savable
	GetID() string
	TypeName() string
	TitleByUser(*User) string
	Posts() []*Post
	PostsRelevantFirst(count int) []*Post
	CountPosts() int
	Creator() *User
	CreatorID() UserID
	AddPost(string)
	RemovePost(string) bool
}
