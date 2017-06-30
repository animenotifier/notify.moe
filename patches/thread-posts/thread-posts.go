package main

import (
	"github.com/animenotifier/arn"
)

func main() {
	// Get a stream of all posts
	allPosts, err := arn.StreamPosts()
	arn.PanicOnError(err)

	threadToPosts := make(map[string][]string)

	// Iterate over the stream
	for post := range allPosts {
		_, found := threadToPosts[post.ThreadID]

		if !found {
			threadToPosts[post.ThreadID] = []string{post.ID}
		} else {
			threadToPosts[post.ThreadID] = append(threadToPosts[post.ThreadID], post.ID)
		}
	}

	// Save new post ID lists
	for threadID, posts := range threadToPosts {
		thread, err := arn.GetThread(threadID)
		arn.PanicOnError(err)

		thread.Posts = posts
		err = thread.Save()
		arn.PanicOnError(err)
	}
}
