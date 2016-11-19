package forum

import (
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

const threadsPerPage = 20

// Get ...
func Get(ctx *aero.Context) string {
	tag := ctx.Get("tag")
	threads, _ := arn.GetThreadsByTag(tag)

	sort.Sort(threads)

	if len(threads) > threadsPerPage {
		threads = threads[:threadsPerPage]
	}

	for _, thread := range threads {
		thread.Init()
	}

	return ctx.HTML(components.Forum(tag, threads))
}
