package paypal

import (
	"net/http"
	"os"

	"github.com/aerogo/aero"
	"github.com/logpacker/PayPal-Go-SDK"
)

// CreatePayment ...
func CreatePayment(ctx *aero.Context) string {
	// Create a client instance
	c, err := paypalsdk.NewClient("clientID", "secretID", paypalsdk.APIBaseSandBox)
	c.SetLog(os.Stdout) // Set log to terminal stdout

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Could not initiate PayPal client", err)
	}

	_, err = c.GetAccessToken()

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Could not get PayPal access token", err)
	}

	amount := paypalsdk.Amount{
		Total:    "7.00",
		Currency: "USD",
	}
	redirectURI := "http://example.com/redirect-uri"
	cancelURI := "http://example.com/cancel-uri"
	description := "Description for this payment"
	paymentResult, err := c.CreateDirectPaypalPayment(amount, redirectURI, cancelURI, description)

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Could not create PayPal payment", err)
	}

	return ctx.JSON(paymentResult)
}
