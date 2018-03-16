package main

import (
	"github.com/animenotifier/arn"
)

func main() {
	defer arn.Node.Close()

	user, _ := arn.GetUser("Vy2Hk5yvx")
	user.Role = ""
	user.Save()
}
