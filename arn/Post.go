package arn

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/aerogo/markdown"
	"github.com/aerogo/nano"
)

// Post is a comment related to any parent type in the database.
type Post struct {
	Tags       []string `json:"tags" editable:"true"`
	ParentID   string   `json:"parentId" editable:"true"`
	ParentType string   `json:"parentType"`
	Edited     string   `json:"edited"`

	hasID
	hasText
	hasPosts
	hasCreator
	hasLikes

	html string
}

// Parent returns the object this post was posted in.
func (post *Post) Parent() PostParent {
	obj, _ := DB.Get(post.ParentType, post.ParentID)
	return obj.(PostParent)
}

// TopMostParent returns the first non-post object this post was posted in.
func (post *Post) TopMostParent() PostParent {
	topMostParent := post.Parent()

	for {
		if topMostParent.TypeName() != "Post" {
			return topMostParent
		}

		newParent := topMostParent.(*Post).Parent()

		if newParent == nil {
			return topMostParent
		}

		topMostParent = newParent
	}
}

// GetParentID returns the object ID of the parent.
func (post *Post) GetParentID() string {
	return post.ParentID
}

// SetParent sets a new parent.
func (post *Post) SetParent(newParent PostParent) {
	// Remove from old parent
	oldParent := post.Parent()
	oldParent.RemovePost(post.ID)
	oldParent.Save()

	// Update own fields
	post.ParentID = newParent.GetID()
	post.ParentType = reflect.TypeOf(newParent).Elem().Name()

	// Add to new parent
	newParent.AddPost(post.ID)
	newParent.Save()
}

// Link returns the relative URL of the post.
func (post *Post) Link() string {
	return "/post/" + post.ID
}

// TypeName returns the type name.
func (post *Post) TypeName() string {
	return "Post"
}

// Self returns the object itself.
func (post *Post) Self() Loggable {
	return post
}

// TitleByUser returns the preferred title for the given user.
func (post *Post) TitleByUser(user *User) string {
	return post.Creator().Nick + "'s comment"
}

// HTML returns the HTML representation of the post.
func (post *Post) HTML() string {
	if post.html != "" {
		return post.html
	}

	post.html = markdown.Render(post.Text)
	return post.html
}

// String implements the default string serialization.
func (post *Post) String() string {
	const maxLen = 170

	if len(post.Text) > maxLen {
		return post.Text[:maxLen-3] + "..."
	}

	return post.Text
}

// OnLike is called when the post receives a like.
func (post *Post) OnLike(likedBy *User) {
	if !post.Creator().Settings().Notification.ForumLikes {
		return
	}

	go func() {
		message := ""
		notifyUser := post.Creator()

		if post.ParentType == "User" {
			if post.ParentID == notifyUser.ID {
				// Somebody liked your post on your own profile
				message = fmt.Sprintf(`%s liked your profile post.`, likedBy.Nick)
			} else {
				// Somebody liked your post on someone else's profile
				message = fmt.Sprintf(`%s liked your post on %s's profile.`, likedBy.Nick, post.Parent().TitleByUser(notifyUser))
			}
		} else {
			message = fmt.Sprintf(`%s liked your post in the %s "%s".`, likedBy.Nick, strings.ToLower(post.ParentType), post.Parent().TitleByUser(notifyUser))
		}

		notifyUser.SendNotification(&PushNotification{
			Title:   likedBy.Nick + " liked your post",
			Message: message,
			Icon:    "https:" + likedBy.AvatarLink("large"),
			Link:    "https://notify.moe" + likedBy.Link(),
			Type:    NotificationTypeLike,
		})
	}()
}

// GetPost ...
func GetPost(id string) (*Post, error) {
	obj, err := DB.Get("Post", id)

	if err != nil {
		return nil, err
	}

	return obj.(*Post), nil
}

// StreamPosts returns a stream of all posts.
func StreamPosts() <-chan *Post {
	channel := make(chan *Post, nano.ChannelBufferSize)

	go func() {
		for obj := range DB.All("Post") {
			channel <- obj.(*Post)
		}

		close(channel)
	}()

	return channel
}

// AllPosts returns a slice of all posts.
func AllPosts() []*Post {
	all := make([]*Post, 0, DB.Collection("Post").Count())

	for obj := range StreamPosts() {
		all = append(all, obj)
	}

	return all
}

// SortPostsLatestFirst sorts the slice of posts.
func SortPostsLatestFirst(posts []*Post) {
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Created > posts[j].Created
	})
}

// SortPostsLatestLast sorts the slice of posts.
func SortPostsLatestLast(posts []*Post) {
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Created < posts[j].Created
	})
}

// FilterPostsWithUniqueThreads removes posts with the same thread until we have enough posts.
func FilterPostsWithUniqueThreads(posts []*Post, limit int) []*Post {
	filtered := []*Post{}
	threadsProcessed := map[string]bool{}

	for _, post := range posts {
		if len(filtered) >= limit {
			return filtered
		}

		_, found := threadsProcessed[post.ParentID]

		if found {
			continue
		}

		threadsProcessed[post.ParentID] = true
		filtered = append(filtered, post)
	}

	return filtered
}

// GetPostsByUser ...
func GetPostsByUser(user *User) ([]*Post, error) {
	var posts []*Post

	for post := range StreamPosts() {
		if post.CreatedBy == user.ID {
			posts = append(posts, post)
		}
	}

	return posts, nil
}

// FilterPosts filters all forum posts by a custom function.
func FilterPosts(filter func(*Post) bool) ([]*Post, error) {
	var filtered []*Post

	for post := range StreamPosts() {
		if filter(post) {
			filtered = append(filtered, post)
		}
	}

	return filtered, nil
}
