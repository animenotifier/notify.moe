package shop

import (
	"net/http"
	"sync"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
)

var itemBuyMutex sync.Mutex

// BuyItem ...
func BuyItem(ctx aero.Context) error {
	// Lock via mutex to prevent race conditions
	itemBuyMutex.Lock()
	defer itemBuyMutex.Unlock()

	// Logged in user
	user := arn.GetUserFromContext(ctx)

	if user == nil {
		return ctx.Error(http.StatusUnauthorized, "Not logged in")
	}

	// Item ID and quantity
	itemID := ctx.Get("item")
	quantity, err := ctx.GetInt("quantity")

	if err != nil || quantity == 0 {
		return ctx.Error(http.StatusBadRequest, "Invalid item quantity", err)
	}

	item, err := arn.GetShopItem(itemID)

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Error fetching item data", err)
	}

	// Calculate total price and subtract balance
	totalPrice := int(item.Price) * quantity

	if user.Balance < totalPrice {
		return ctx.Error(http.StatusBadRequest, "Not enough gems. You need to charge up your balance before you can buy this item.")
	}

	// Add item to user inventory
	inventory := user.Inventory()
	err = inventory.AddItem(itemID, uint(quantity))

	if err != nil {
		return ctx.Error(http.StatusBadRequest, err)
	}

	inventory.Save()

	// Deduct balance
	user.Balance -= totalPrice
	user.Save()

	// Save purchase
	purchase := arn.NewPurchase(user.ID, itemID, quantity, int(item.Price), "gem")
	purchase.Save()

	return nil
}
