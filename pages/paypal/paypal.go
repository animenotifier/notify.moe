package paypal

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/assets"
	paypalsdk "github.com/logpacker/PayPal-Go-SDK"
)

// CreatePayment creates the PayPal payment, typically via a JSON API route.
func CreatePayment(ctx aero.Context) error {
	// Make sure the user is logged in
	user := arn.GetUserFromContext(ctx)

	if user == nil {
		return ctx.Error(http.StatusUnauthorized, "Not logged in")
	}

	// Verify amount
	amount, err := ctx.Request().Body().String()

	if err != nil {
		return ctx.Error(http.StatusBadRequest, "Could not read amount", err)
	}

	switch amount {
	case "1000", "2000", "3000", "6000", "12000", "25000", "50000", "75000":
		// OK
	default:
		return ctx.Error(http.StatusBadRequest, "Incorrect amount")
	}

	// Initiate PayPal client
	c, err := arn.PayPal()

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Could not initiate PayPal client", err)
	}

	// Get access token
	_, err = c.GetAccessToken()

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Could not get PayPal access token", err)
	}

	// Create payment
	p := paypalsdk.Payment{
		Intent: "sale",
		Payer: &paypalsdk.Payer{
			PaymentMethod: "paypal",
		},
		Transactions: []paypalsdk.Transaction{{
			Amount: &paypalsdk.Amount{
				Currency: "JPY",
				Total:    amount,
			},
			Description: "Top Up Balance",
		}},
		RedirectURLs: &paypalsdk.RedirectURLs{
			ReturnURL: "https://" + assets.Domain + "/paypal/success",
			CancelURL: "https://" + assets.Domain + "/paypal/cancel",
		},
	}

	paymentResponse, err := c.CreatePayment(p)

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Could not create PayPal payment", err)
	}

	return ctx.JSON(paymentResponse)
}
