package editpost

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
	"github.com/animenotifier/notify.moe/utils/editform"
)

// Get post edit page.
func Get(ctx *aero.Context) string {
	id := ctx.Get("id")
	user := utils.GetUser(ctx)

	if user == nil || (user.Role != "editor" && user.Role != "admin") {
		return ctx.Error(http.StatusUnauthorized, "Not logged in or not auhorized to edit this post")
	}

	post, err := arn.GetPost(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Post not found", err)
	}

	return ctx.HTML(components.EditPostTabs(post) + editform.Render(post, "Edit post", user))
}
