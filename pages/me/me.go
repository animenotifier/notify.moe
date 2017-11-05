package me

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/utils"
)

// Get ...
func Get(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	if user == nil {
		return ctx.JSON(nil)
	}

	return ctx.JSON(user)
}
