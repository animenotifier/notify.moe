package paypal

// Web profile is currently disabled.

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
