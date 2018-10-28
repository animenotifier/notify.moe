package main

import (
	"github.com/animenotifier/arn"
	"github.com/animenotifier/arn/stringutils"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Updating AMV info")

	defer color.Green("Finished.")
	defer arn.Node.Close()

	for amv := range arn.StreamAMVs() {
		err := amv.RefreshInfo()

		if err != nil {
			color.Red(err.Error())
			continue
		}

		color.Green(amv.Title.Canonical)
		stringutils.PrettyPrint(amv.Info)
		amv.Save()
	}
}
