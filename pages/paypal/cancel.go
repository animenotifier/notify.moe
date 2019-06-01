package paypal

import (
	"fmt"

	"github.com/aerogo/aero"
)

// Cancel ...
func Cancel(ctx aero.Context) error {
	token := ctx.Query("token")
	fmt.Println("cancel", token)

	return ctx.HTML("Payment has been canceled.")
}
