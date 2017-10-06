package shop

import (
	"net/http"
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// PurchaseHistory ...
func PurchaseHistory(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	if user == nil {
		return ctx.Error(http.StatusUnauthorized, "Not logged in", nil)
	}

	purchases, err := arn.AllPurchases()

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Error fetching shop item data", err)
	}

	sort.Slice(purchases, func(i, j int) bool {
		return purchases[i].Date > purchases[j].Date
	})

	return ctx.HTML(components.PurchaseHistory(purchases, user))
}
