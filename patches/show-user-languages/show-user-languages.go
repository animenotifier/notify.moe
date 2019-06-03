package main

import (
	"fmt"

	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	languages := map[string]int{}

	for user := range arn.StreamUsers() {
		languages[user.Language]++
	}

	for language, users := range languages {
		fmt.Printf("* [%s] %d users\n", language, users)
	}
}
