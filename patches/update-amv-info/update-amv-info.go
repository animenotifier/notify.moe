package main

import (
	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/arn/stringutils"
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
