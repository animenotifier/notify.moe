package main

import (
	"flag"

	"github.com/animenotifier/arn"
)

// Shell parameters
var userID string
var userNick string

// Shell flags
func init() {
	flag.StringVar(&userID, "id", "", "ID of the user whose avatar you want to refresh")
	flag.StringVar(&userNick, "nick", "", "Nickname of the user whose avatar you want to refresh")
	flag.Parse()
}

// InvokeShellArgs ...
func InvokeShellArgs() bool {
	if userID != "" {
		user, err := arn.GetUser(userID)

		if err != nil {
			panic(err)
		}

		Work(user)
		return true
	}

	if userNick != "" {
		user, err := arn.GetUserByNick(userNick)

		if err != nil {
			panic(err)
		}

		Work(user)
		return true
	}

	return false
}
