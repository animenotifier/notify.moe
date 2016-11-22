package threads

import (
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

// Get ...
func Get(ctx *aero.Context) string {
	id := ctx.Get("id")
	thread, err := arn.GetThread(id)

	if err != nil {
		return ctx.Error(404, "Thread not found")
	}

	replies, filterErr := arn.FilterPosts(func(post *arn.Post) bool {
		return post.ThreadID == thread.ID
	})

	sort.Sort(replies)

	if filterErr != nil {
		return ctx.Error(500, "Error fetching thread replies")
	}

	return ctx.HTML(components.Thread(thread, replies))
}
