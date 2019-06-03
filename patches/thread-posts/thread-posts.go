package main

import (
	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	defer arn.Node.Close()

	// Get a stream of all posts
	threadToPosts := make(map[string][]string)

	// Iterate over the stream
	for post := range arn.StreamPosts() {
		if post.ParentType != "Thread" {
			continue
		}

		_, found := threadToPosts[post.ParentID]

		if !found {
			threadToPosts[post.ParentID] = []string{post.ID}
		} else {
			threadToPosts[post.ParentID] = append(threadToPosts[post.ParentID], post.ID)
		}
	}

	// Save new post ID lists
	for threadID, posts := range threadToPosts {
		thread, err := arn.GetThread(threadID)
		arn.PanicOnError(err)

		thread.PostIDs = posts
		thread.Save()
	}
}
