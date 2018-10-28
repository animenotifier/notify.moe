package main

import (
	"path"

	"github.com/animenotifier/arn"
	"github.com/animenotifier/arn/stringutils"
	"github.com/animenotifier/arn/video"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Updating AMV info")

	defer color.Green("Finished.")
	defer arn.Node.Close()

	for amv := range arn.StreamAMVs() {
		if amv.File == "" {
			continue
		}

		info, err := video.GetInfo(path.Join(arn.Root, "videos", "amvs", amv.File))

		if err != nil {
			color.Red(err.Error())
			continue
		}

		color.Green(amv.Title.Canonical)
		stringutils.PrettyPrint(info)

		amv.Info = *info
		amv.Save()
	}
}
