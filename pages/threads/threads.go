package threads

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Get thread.
func Get(ctx *aero.Context) string {
	id := ctx.Get("id")
	user := utils.GetUser(ctx)

	// Fetch thread
	thread, err := arn.GetThread(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Thread not found", err)
	}

	// Fetch posts
	postObjects, getErr := arn.DB.GetMany("Post", thread.Posts)

	if getErr != nil {
		return ctx.Error(http.StatusInternalServerError, "Could not retrieve posts", getErr)
	}

	posts := postObjects.([]*arn.Post)

	// Sort posts
	arn.SortPostsLatestLast(posts)

	return ctx.HTML(components.Thread(thread, posts, user))
}
