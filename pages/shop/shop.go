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
func Get(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	if user == nil {
		return ctx.Error(http.StatusUnauthorized, "Not logged in", nil)
	}

	items, err := arn.AllShopItems()

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Error fetching shop item data", err)
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].Order < items[j].Order
	})

	return ctx.HTML(components.Shop(user, items))
}
