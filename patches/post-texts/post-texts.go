package main

import (
	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/arn/autocorrect"
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
