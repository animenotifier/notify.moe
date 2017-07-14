package notifications

import (
	"fmt"
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/utils"
)

// Test ...
func Test(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	if user == nil {
		return ctx.Error(http.StatusBadRequest, "Not logged in", nil)
	}

	for _, sub := range user.PushSubscriptions().Items {
		err := sub.SendNotification("Yay, it works!")

		if err != nil {
			fmt.Println(err)
		}
	}

	return "ok"
}
