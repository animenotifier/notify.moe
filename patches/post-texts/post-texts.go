package main

import (
	"github.com/animenotifier/arn"
	"github.com/animenotifier/arn/autocorrect"
	"github.com/fatih/color"
)

func main() {
	defer arn.Node.Close()

	// Iterate over the stream
	for post := range arn.StreamPosts() {
		// Fix text
		color.Yellow(post.Text)
		post.Text = autocorrect.PostText(post.Text)
		color.Green(post.Text)

		// Tags
		if post.Tags == nil {
			post.Tags = []string{}
		}

		// Save
		post.Save()
	}
}
