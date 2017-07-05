package utils

import "github.com/aerogo/aero"

// AllowEmbed allows the page to be called by the browser extension.
func AllowEmbed(ctx *aero.Context, response string) string {
	// This is a bit of a hack.
	ctx.SetResponseHeader("X-Frame-Options", "ALLOW-FROM chrome-extension://hjfcooigdelogjmniiahfiilcefdlpha/options.html")
	return response
}
