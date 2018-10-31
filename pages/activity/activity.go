package activity

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

const maxActivitiesPerPage = 40

// Get activity page.
func Get(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	// posts := arn.AllPosts()
	// arn.SortPostsLatestFirst(posts)

	// posts := arn.FilterPosts(func(post *arn.Post) bool {
	// 	return post.
	// })

	entries := arn.FilterEditLogEntries(func(entry *arn.EditLogEntry) bool {
		return entry.Action == "create" && entry.ObjectType == "Post" && entry.Object() != nil
	})

	arn.SortEditLogEntriesLatestFirst(entries)

	if len(entries) > maxActivitiesPerPage {
		entries = entries[:maxActivitiesPerPage]
	}

	return ctx.HTML(components.ActivityFeed(entries, user))
}
