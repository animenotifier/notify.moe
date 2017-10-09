package main

import (
	"fmt"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Addind draft indices")

	// Iterate over the stream
	for user := range arn.MustStreamUsers() {
		fmt.Println(user.Nick)

		draftIndex := arn.NewDraftIndex(user.ID)
		arn.PanicOnError(draftIndex.Save())
	}

	color.Green("Finished.")
}
