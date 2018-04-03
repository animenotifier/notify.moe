package home

import (
	"strings"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/pages/frontpage"
	"github.com/animenotifier/notify.moe/utils"
)

// Get the anime list or the frontpage when logged out.
func Get(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	if user == nil {
		return frontpage.Get(ctx)
	}

	// Redirect
	prefix := "/"

	if strings.HasPrefix(ctx.URI(), "/_") {
		prefix = "/_/"
	}

	return ctx.Redirect(prefix + "+" + user.Nick + "/animelist/watching")
}
