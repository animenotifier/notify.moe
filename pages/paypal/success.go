package paypal

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
)

// Success ...
func Success(ctx *aero.Context) string {
	paymentID := ctx.Query("paymentId")
	// token := ctx.Query("token")
	// payerID := ctx.Query("PayerID")

	c, err := arn.PayPal()

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Could not initiate PayPal client", err)
	}

	_, err = c.GetAccessToken()

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Could not get PayPal access token", err)
	}

	payment, err := c.GetPayment(paymentID)

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Could not retrieve payment information", err)
	}

	arn.PrettyPrint(payment)

	return ctx.HTML("success")
}
