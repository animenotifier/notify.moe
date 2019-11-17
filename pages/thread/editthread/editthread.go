package editthread

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils/editform"
)

// Get thread edit page.
func Get(ctx aero.Context) error {
	id := ctx.Get("id")
	user := arn.GetUserFromContext(ctx)

	if user == nil || (user.Role != "editor" && user.Role != "admin") {
		return ctx.Error(http.StatusUnauthorized, "Not logged in or not auhorized to edit this thread")
	}

	thread, err := arn.GetThread(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Thread not found", err)
	}

	return ctx.HTML(components.EditThreadTabs(thread) + editform.Render(thread, "Edit thread", user))
}
