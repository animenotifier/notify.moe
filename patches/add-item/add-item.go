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
	flag.StringVar(&nick, "nick", "", "Name of the user.")
	flag.StringVar(&itemID, "item", "", "ID of the item.")
	flag.IntVar(&quantity, "q", 1, "Item quantity.")
	flag.Parse()
}

func main() {
	defer arn.Node.Close()

	if nick == "" || itemID == "" {
		color.Red("Missing parameters")
		return
	}

	user, err := arn.GetUserByNick(nick)
	arn.PanicOnError(err)

	item, err := arn.GetShopItem(itemID)
	arn.PanicOnError(err)

	if item == nil {
		color.Red("Unknown item")
		return
	}

	// Add to user inventory
	inventory := user.Inventory()
	inventory.AddItem(itemID, uint(quantity))
	inventory.Save()
}
