package main

import (
	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	// Get a stream of all posts
	allPosts, err := arn.AllPosts()
	arn.PanicOnError(err)

	// Iterate over the stream
	for post := range allPosts {
		// Fix text
		color.Yellow(post.Text)
		post.Text = arn.FixPostText(post.Text)
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
