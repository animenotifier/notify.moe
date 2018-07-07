package main

import (
	"fmt"
	"time"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

const month = 30 * 24 * time.Hour

func main() {
	color.Yellow("Deleting private user data")

	defer arn.Node.Close()
	defer color.Green("Finished.")

	for user := range arn.StreamUsers() {
		color.Cyan(user.Nick)

		for _, notification := range user.Notifications().Notifications() {
			if time.Since(notification.CreatedTime()) < 2*month {
				continue
			}

			fmt.Println(notification)
			// notification.Delete()
		}

		// Save in DB
		// user.Save()
	}
}
