package main

import (
	"fmt"

	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	defer arn.Node.Close()

	for user := range arn.StreamUsers() {
		if user.Notifications().CountUnseen() <= 10 && !user.IsActive() && user.Email != "" && len(user.AnimeList().Items) == 0 {
			fmt.Println(user.Email)
		}
	}
}
