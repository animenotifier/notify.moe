package main

import (
	"flag"

	"github.com/animenotifier/notify.moe/arn"
)

// Shell parameters
var userID string
var role string

// Shell flags
func init() {
	flag.StringVar(&userID, "id", "", "ID of the user")
	flag.StringVar(&role, "role", "", "The user's new role")
	flag.Parse()
}

func main() {
	defer arn.Node.Close()

	// Show usage if needed
	if userID == "" || role == "" {
		flag.Usage()
		return
	}

	// Get user
	user, err := arn.GetUser(userID)
	arn.PanicOnError(err)

	// Save role
	user.Role = role
	user.Save()
}
