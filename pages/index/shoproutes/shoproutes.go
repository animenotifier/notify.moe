package shoproutes

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/pages/charge"
	"github.com/animenotifier/notify.moe/pages/inventory"
	"github.com/animenotifier/notify.moe/pages/paypal"
	"github.com/animenotifier/notify.moe/pages/shop"
	"github.com/animenotifier/notify.moe/pages/support"
	"github.com/animenotifier/notify.moe/utils/page"
)

// Register registers the page routes.
func Register(app *aero.Application) {
	// Shop
	page.Get(app, "/support", support.Get)
	page.Get(app, "/shop", shop.Get)
	page.Get(app, "/inventory", inventory.Get)
	page.Get(app, "/charge", charge.Get)
	page.Get(app, "/shop/history", shop.PurchaseHistory)

	// PayPal
	page.Get(app, "/paypal/success", paypal.Success)
	page.Get(app, "/paypal/cancel", paypal.Cancel)

	// API: Create payment
	app.Post("/api/paypal/payment/create", paypal.CreatePayment)

	// API: Buy item
	app.Post("/api/shop/buy/:item/:quantity", shop.BuyItem)
}
