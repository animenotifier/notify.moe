package main

import (
	"flag"

	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
)

// Shell parameters
var nick string
var service string
var serviceID string

// Shell flags
func init() {
	flag.StringVar(&nick, "nick", "", "Nick of the user")
	flag.StringVar(&service, "service", "", "Service name (Google or Facebook)")
	flag.StringVar(&serviceID, "serviceID", "", "ID of the user on the given service")
	flag.Parse()
}

func main() {
	if nick == "" || service == "" || serviceID == "" {
		flag.Usage()
		return
	}

	color.Yellow("Updating user service ID")

	defer color.Green("Finished.")
	defer arn.Node.Close()

	user, err := arn.GetUserByNick(nick)
	arn.PanicOnError(err)

	switch service {
	case "Google":
		user.ConnectGoogle(serviceID)

	case "Facebook":
		user.ConnectFacebook(serviceID)

	default:
		panic("Unknown service")
	}

	user.Save()
}
