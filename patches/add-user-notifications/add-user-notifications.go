package main

import (
	"fmt"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	defer arn.Node.Close()

	for user := range arn.StreamUsers() {
		obj := user.Notifications()

		if obj == nil {
			fmt.Println(user.Nick)
			arn.NewUserNotifications(user.ID).Save()
		}
	}

	color.Green("Finished.")
}
