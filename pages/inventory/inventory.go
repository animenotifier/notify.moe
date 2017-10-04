package inventory

import (
	"net/http"

	"github.com/animenotifier/arn"

	"github.com/animenotifier/notify.moe/components"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/utils"
)

// Get inventory page.
func Get(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	if user == nil {
		return ctx.Error(http.StatusUnauthorized, "Not logged in", nil)
	}

	inventory, err := arn.GetInventory(user.ID)

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Error fetching inventory data", err)
	}

	// TEST
	inventory.AddItem("anime-support-ticket", 35)
	inventory.AddItem("pro-account-24", 20)
	inventory.AddItem("anime-support-ticket", 15)
	inventory.AddItem("pro-account-24", 10)

	return ctx.HTML(components.Inventory(inventory))
}
