package main

import (
	"github.com/animenotifier/arn"
)

func main() {
	defer arn.Node.Close()

	user, _ := arn.GetUserByNick("Akyoto")

	if user.Language == ":)" {
		user.Language = ":("
	} else {
		user.Language = ":)"
	}

	user.Save()
}
