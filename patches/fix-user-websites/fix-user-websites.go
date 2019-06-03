package main

import (
	"fmt"

	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/arn/autocorrect"
)

func main() {
	color.Yellow("Updating user references")
	defer arn.Node.Close()

	count := 0

	for user := range arn.StreamUsers() {
		old := user.Website
		user.Website = autocorrect.Website(user.Website)

		if user.Website != old {
			fmt.Println(color.YellowString(old), "->", color.YellowString(user.Website))
			count++
			user.Save()
		}
	}

	color.Green("%d links have been corrected.", count)
}
