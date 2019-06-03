package arn

import (
	"sort"
)

// HasPosts includes a list of Post IDs.
type hasPosts struct {
	PostIDs []string `json:"posts"`
}

// AddPost adds a post to the object.
func (obj *hasPosts) AddPost(postID string) {
	obj.PostIDs = append(obj.PostIDs, postID)
}

// RemovePost removes a post from the object.
func (obj *hasPosts) RemovePost(postID string) bool {
	for index, item := range obj.PostIDs {
		if item == postID {
			obj.PostIDs = append(obj.PostIDs[:index], obj.PostIDs[index+1:]...)
			return true
		}
	}

	return false
}

// Posts returns a slice of all posts.
func (obj *hasPosts) Posts() []*Post {
	objects := DB.GetMany("Post", obj.PostIDs)
	posts := make([]*Post, 0, len(objects))

	for _, post := range objects {
		if post == nil {
			continue
		}

		posts = append(posts, post.(*Post))
	}

	return posts
}

// PostsRelevantFirst returns a slice of all posts sorted by relevance.
func (obj *hasPosts) PostsRelevantFirst(count int) []*Post {
	original := obj.Posts()
	newPosts := make([]*Post, len(original))
	copy(newPosts, original)

	sort.Slice(newPosts, func(i, j int) bool {
		return newPosts[i].Created > newPosts[j].Created
	})

	if count >= 0 && len(newPosts) > count {
		newPosts = newPosts[:count]
	}

	return newPosts
}

// CountPosts returns the number of posts written for this object.
func (obj *hasPosts) CountPosts() int {
	return len(obj.PostIDs)
}
