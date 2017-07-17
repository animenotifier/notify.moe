package main

import (
	"strings"

	"github.com/animenotifier/arn"
)

func main() {
	for user := range arn.MustStreamUsers() {
		if !strings.HasPrefix(user.Avatar.Extension, ".") {
			user.Avatar.Extension = ""
		}

		user.Save()
	}
}
