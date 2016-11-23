package profile

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

// Get ...
func Get(ctx *aero.Context) string {
	nick := ctx.Get("nick")
	user, err := arn.GetUserByNick(nick)

	if err != nil {
		return ctx.Error(404, "User not found", err)
	}

	return ctx.HTML(components.Profile(user, nil))
}
