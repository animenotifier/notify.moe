package admin

import (
	"net/http"
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
)

// PurchaseHistory ...
func PurchaseHistory(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)

	if user == nil {
		return ctx.Error(http.StatusUnauthorized, "Not logged in")
	}

	if user.Role != "admin" {
		return ctx.Error(http.StatusUnauthorized, "Not authorized")
	}

	purchases, err := arn.AllPurchases()

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Error fetching shop item data", err)
	}

	sort.Slice(purchases, func(i, j int) bool {
		return purchases[i].Date > purchases[j].Date
	})

	return ctx.HTML(components.GlobalPurchaseHistory(purchases))
}
