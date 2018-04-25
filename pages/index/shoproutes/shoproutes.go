package shoproutes

import (
	"github.com/aerogo/aero"
	"github.com/aerogo/layout"
	"github.com/animenotifier/notify.moe/pages/charge"
	"github.com/animenotifier/notify.moe/pages/inventory"
	"github.com/animenotifier/notify.moe/pages/paypal"
	"github.com/animenotifier/notify.moe/pages/shop"
	"github.com/animenotifier/notify.moe/pages/support"
)

// Register registers the page routes.
func Register(l *layout.Layout, app *aero.Application) {
	// Shop
	l.Page("/support", support.Get)
	l.Page("/shop", shop.Get)
	l.Page("/inventory", inventory.Get)
	l.Page("/charge", charge.Get)
	l.Page("/shop/history", shop.PurchaseHistory)

	// PayPal
	l.Page("/paypal/success", paypal.Success)
	l.Page("/paypal/cancel", paypal.Cancel)

	// API: Create payment
	app.Post("/api/paypal/payment/create", paypal.CreatePayment)

	// API: Buy item
	app.Post("/api/shop/buy/:item/:quantity", shop.BuyItem)
}
