package main

import (
	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	defer arn.Node.Close()

	for user := range arn.StreamUsers() {
		settings := user.Settings()

		if !user.IsPro() {
			settings.Theme = "light"
		}

		settings.Save()
	}
}
