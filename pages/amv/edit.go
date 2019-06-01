package amv

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
	"github.com/animenotifier/notify.moe/utils/editform"
)

// Edit track.
func Edit(ctx aero.Context) error {
	id := ctx.Get("id")
	amv, err := arn.GetAMV(id)
	user := utils.GetUser(ctx)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "AMV not found", err)
	}

	return ctx.HTML(components.AMVTabs(amv, user) + editform.Render(amv, "Edit AMV", user))
}
