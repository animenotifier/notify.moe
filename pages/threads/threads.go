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
		return ctx.Text("Thread not found")
	}

	thread.Author, _ = arn.GetUser(thread.AuthorID)

	replies, filterErr := arn.FilterPosts(func(post *arn.Post) bool {
		return post.ThreadID == thread.ID
	})

	sort.Sort(replies)

	if filterErr != nil {
		return ctx.Text("Error fetching thread replies")
	}

	return ctx.HTML(components.Thread(thread, replies))
}
