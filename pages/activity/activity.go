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
		if entry.Action != "create" {
			return false
		}

		obj := entry.Object()

		if obj == nil {
			return false
		}

		_, isPostable := obj.(arn.Postable)
		return isPostable
	})

	arn.SortEditLogEntriesLatestFirst(entries)

	if len(entries) > maxActivitiesPerPage {
		entries = entries[:maxActivitiesPerPage]
	}

	return ctx.HTML(components.ActivityFeed(entries, user))
}
