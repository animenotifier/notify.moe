package forumroutes

import (
	"github.com/aerogo/layout"
	"github.com/animenotifier/notify.moe/pages/forum"
	"github.com/animenotifier/notify.moe/pages/newthread"
	"github.com/animenotifier/notify.moe/pages/post"
	"github.com/animenotifier/notify.moe/pages/post/editpost"
	"github.com/animenotifier/notify.moe/pages/thread"
	"github.com/animenotifier/notify.moe/pages/thread/editthread"
)

// Register registers the page routes.
func Register(l *layout.Layout) {
	// Forum
	l.Page("/forum", forum.Get)
	l.Page("/forum/:tag", forum.Get)

	// Thread
	l.Page("/thread/:id", thread.Get)
	l.Page("/thread/:id/edit", editthread.Get)
	l.Page("/new/thread", newthread.Get)

	// Post
	l.Page("/post/:id", post.Get)
	l.Page("/post/:id/edit", editpost.Get)
}
