package main

import (
	"github.com/animenotifier/arn"
)

func main() {
	defer arn.Node.Close()

	// Get a stream of all posts
	threadToPosts := make(map[string][]string)

	// Iterate over the stream
	for post := range arn.StreamPosts() {
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
		thread.Save()
	}
}
