package main

import (
	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	defer arn.Node.Close()

	for follows := range arn.StreamUserFollows() {
		user, err := arn.GetUser(follows.UserID)

		if err != nil {
			color.Red(err.Error())
			continue
		}

		user.FollowIDs = follows.Items
		user.Save()
	}
}
