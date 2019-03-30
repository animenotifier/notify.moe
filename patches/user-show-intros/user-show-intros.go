package main

import (
	"regexp"

	"github.com/animenotifier/arn"
	"github.com/blitzprog/color"
)

var flaggedWords = regexp.MustCompile("fuck|fucking|freaking|shit|bad|terrible|awful|wtf")

func main() {
	color.Yellow("Showing user intros")

	defer color.Green("Finished.")
	defer arn.Node.Close()

	for user := range arn.StreamUsers() {
		if user.Introduction == "" {
			continue
		}

		if flaggedWords.MatchString(user.Introduction) {
			color.Cyan(user.Nick)
			color.Red(user.Introduction)
		}
	}
}
