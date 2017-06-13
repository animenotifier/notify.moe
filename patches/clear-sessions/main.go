package main

import (
	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Deleting all sessions...")
	arn.DB.DeleteTable("Session")
	color.Green("Finished.")
}
