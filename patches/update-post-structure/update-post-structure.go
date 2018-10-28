package main

import (
	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Updating post structure")

	defer color.Green("Finished.")
	defer arn.Node.Close()

	// Iterate over the stream
	for post := range arn.StreamPosts() {
		post.ParentID = post.ThreadID
		post.ParentType = "Thread"
		post.Save()
	}
}
