package main

import (
	"flag"

	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
)

var nick string
var itemID string
var quantity int

func init() {
	flag.StringVar(&nick, "nick", "", "Name of the user. Leave it out to target all users.")
	flag.StringVar(&itemID, "item", "", "ID of the item.")
	flag.IntVar(&quantity, "q", 1, "Item quantity.")
	flag.Parse()
}

func main() {
	defer arn.Node.Close()

	if itemID == "" {
		flag.Usage()
		return
	}

	// Check that the item exists
	_, err := arn.GetShopItem(itemID)
	arn.PanicOnError(err)

	if nick != "" {
		// Single user
		user, err := arn.GetUserByNick(nick)
		arn.PanicOnError(err)
		err = addItemToUser(user)
		arn.PanicOnError(err)
	} else {
		// All users
		for user := range arn.StreamUsers() {
			err = addItemToUser(user)

			if err != nil {
				color.Red(err.Error())
			}
		}
	}
}

func addItemToUser(user *arn.User) error {
	inventory := user.Inventory()
	err := inventory.AddItem(itemID, uint(quantity))

	if err != nil {
		return err
	}

	inventory.Save()
	return nil
}
