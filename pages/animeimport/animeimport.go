package animeimport

import (
	"net/http"

	"github.com/aerogo/aero"
)

// Kitsu anime import.
func Kitsu(ctx *aero.Context) string {
	// id := ctx.Get("id")
	// user := utils.GetUser(ctx)

	if true {
		return ctx.Error(http.StatusUnauthorized, "Not authorized", nil)
	}

	return ""
}
