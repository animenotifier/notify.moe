package forum

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

const threadsPerPage = 20

// Get forum category.
func Get(ctx *aero.Context) string {
	tag := ctx.Get("tag")
	threads, _ := arn.GetThreadsByTag(tag)
	arn.SortThreads(threads)

	if len(threads) > threadsPerPage {
		threads = threads[:threadsPerPage]
	}

	return ctx.HTML(components.Forum(tag, threads))
}
