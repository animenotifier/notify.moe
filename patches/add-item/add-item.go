package main

import (
	"flag"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
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
		color.Red("Missing parameters")
		return
	}

	// Check that the item exists
	_, err := arn.GetShopItem(itemID)
	arn.PanicOnError(err)

	if nick != "" {
		// Single user
		user, err := arn.GetUserByNick(nick)
		arn.PanicOnError(err)
		addItemToUser(user)
	} else {
		// All users
		for user := range arn.StreamUsers() {
			addItemToUser(user)
		}
	}
}

func addItemToUser(user *arn.User) {
	inventory := user.Inventory()
	inventory.AddItem(itemID, uint(quantity))
	inventory.Save()
}
