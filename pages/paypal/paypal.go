package paypal

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/utils"
	"github.com/logpacker/PayPal-Go-SDK"
)

// CreatePayment ...
func CreatePayment(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	if user == nil {
		return ctx.Error(http.StatusUnauthorized, "Not logged in", nil)
	}

	amount, err := ctx.Request().Body().String()

	if err != nil {
		return ctx.Error(http.StatusBadRequest, "Could not read amount", err)
	}

	// Verify amount
	switch amount {
	case "1000", "2000", "3000", "6000", "12000":
		// OK
	default:
		return ctx.Error(http.StatusBadRequest, "Incorrect amount", nil)
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

	// webprofile := paypalsdk.WebProfile{
	// 	Name: "Anime Notifier",
	// 	Presentation: paypalsdk.Presentation{
	// 		BrandName:  "Anime Notifier",
	// 		LogoImage:  "https://notify.moe/brand/220",
	// 		LocaleCode: "US",
	// 	},

	// 	InputFields: paypalsdk.InputFields{
	// 		AllowNote:       true,
	// 		NoShipping:      paypalsdk.NoShippingDisplay,
	// 		AddressOverride: paypalsdk.AddrOverrideFromCall,
	// 	},

	// 	FlowConfig: paypalsdk.FlowConfig{
	// 		LandingPageType: paypalsdk.LandingPageTypeBilling,
	// 	},
	// }

	// result, err := c.CreateWebProfile(webprofile)
	// c.SetWebProfile(*result)

	// if err != nil {
	// 	return ctx.Error(http.StatusInternalServerError, "Could not create PayPal web profile", err)
	// }

	// total := amount[:len(amount)-2] + "." + amount[len(amount)-2:]

	// Create payment
	p := paypalsdk.Payment{
		Intent: "sale",
		Payer: &paypalsdk.Payer{
			PaymentMethod: "paypal",
		},
		Transactions: []paypalsdk.Transaction{paypalsdk.Transaction{
			Amount: &paypalsdk.Amount{
				Currency: "JPY",
				Total:    amount,
			},
			Description: "Top Up Balance",
		}},
		RedirectURLs: &paypalsdk.RedirectURLs{
			ReturnURL: "https://" + ctx.App.Config.Domain + "/paypal/success",
			CancelURL: "https://" + ctx.App.Config.Domain + "/paypal/cancel",
		},
	}

	paymentResponse, err := c.CreatePayment(p)

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Could not create PayPal payment", err)
	}

	return ctx.JSON(paymentResponse)
}
