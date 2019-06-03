package forumroutes

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/pages/forum"
	"github.com/animenotifier/notify.moe/pages/newthread"
	"github.com/animenotifier/notify.moe/pages/post"
	"github.com/animenotifier/notify.moe/pages/post/editpost"
	"github.com/animenotifier/notify.moe/pages/thread"
	"github.com/animenotifier/notify.moe/pages/thread/editthread"
	"github.com/animenotifier/notify.moe/utils/page"
)

// Register registers the page routes.
func Register(app *aero.Application) {
	// Forum
	page.Get(app, "/forum", forum.Get)
	page.Get(app, "/forum/:tag", forum.Get)

	// Thread
	page.Get(app, "/thread/:id", thread.Get)
	page.Get(app, "/thread/:id/edit", editthread.Get)
	page.Get(app, "/new/thread", newthread.Get)

	// Post
	page.Get(app, "/post/:id", post.Get)
	page.Get(app, "/post/:id/edit", editpost.Get)
}
