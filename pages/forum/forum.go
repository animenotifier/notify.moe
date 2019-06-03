package forum

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
)

// ThreadsPerPage indicates how many threads are shown on one page.
const ThreadsPerPage = 20

// Get forum category.
func Get(ctx aero.Context) error {
	tag := ctx.Get("tag")
	threads := arn.GetThreadsByTag(tag)
	arn.SortThreads(threads)

	if len(threads) > ThreadsPerPage {
		threads = threads[:ThreadsPerPage]
	}

	return ctx.HTML(components.Forum(tag, threads, ThreadsPerPage))
}
