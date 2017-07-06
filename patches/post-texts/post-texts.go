package main

import (
	"github.com/animenotifier/arn"
	"github.com/animenotifier/arn/autocorrect"
	"github.com/fatih/color"
)

func main() {
	// Get a stream of all posts
	allPosts, err := arn.StreamPosts()
	arn.PanicOnError(err)

	// Iterate over the stream
	for post := range allPosts {
		// Fix text
		color.Yellow(post.Text)
		post.Text = autocorrect.FixPostText(post.Text)
		color.Green(post.Text)

		// Tags
		if post.Tags == nil {
			post.Tags = []string{}
		}

		// Save
		err = post.Save()
		arn.PanicOnError(err)
	}
}
