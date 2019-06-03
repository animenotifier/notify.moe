package main

import (
	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	defer arn.Node.Close()

	color.Yellow("Deleting all sessions...")
	arn.DB.Clear("Session")
	color.Green("Finished.")
}
