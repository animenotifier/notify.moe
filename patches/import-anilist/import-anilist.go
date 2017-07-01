package main

import (
	"github.com/animenotifier/arn"
)

func main() {
	arn.PanicOnError(arn.AniList.Authorize())
}
