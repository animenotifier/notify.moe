package shop

import (
	"net/http"
	"sort"

	"github.com/animenotifier/arn"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Get shop page.
func Get(ctx aero.Context) error {
	user := utils.GetUser(ctx)

	if user == nil {
		return ctx.Error(http.StatusUnauthorized, "Not logged in")
	}

	items, err := arn.AllShopItems()

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Error fetching shop item data", err)
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].Order < items[j].Order
	})

	// Calculate popularity of items
	itemPopularity := map[string]int{}

	for purchase := range arn.StreamPurchases() {
		itemPopularity[purchase.ItemID]++
	}

	return ctx.HTML(components.Shop(items, itemPopularity, user))
}
