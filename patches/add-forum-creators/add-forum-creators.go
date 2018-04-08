package main

import "github.com/animenotifier/arn"

func main() {
	defer arn.Node.Close()

	for post := range arn.StreamPosts() {
		post.CreatedBy = post.AuthorID
		post.Save()
	}

	for thread := range arn.StreamThreads() {
		thread.CreatedBy = thread.AuthorID
		thread.Save()
	}
}
