package admin

import (
	"net/http"
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
)

// PaymentHistory ...
func PaymentHistory(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)

	if user == nil {
		return ctx.Error(http.StatusUnauthorized, "Not logged in")
	}

	if user.Role != "admin" {
		return ctx.Error(http.StatusUnauthorized, "Not authorized")
	}

	payments, err := arn.AllPayPalPayments()

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Error fetching shop item data", err)
	}

	sort.Slice(payments, func(i, j int) bool {
		return payments[i].Created > payments[j].Created
	})

	return ctx.HTML(components.GlobalPaymentHistory(payments))
}
