package main

import (
	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	defer arn.Node.Close()

	for post := range arn.StreamPosts() {
		if post.Parent() == nil {
			color.Yellow(post.ID)
			color.Red(post.Text)
		}
	}
}
