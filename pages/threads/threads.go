package threads

import (
	"strings"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Get thread.
func Get(ctx *aero.Context) string {
	id := ctx.Get("id")
	thread, err := arn.GetThread(id)
	user := utils.GetUser(ctx)

	if err != nil {
		return ctx.Error(404, "Thread not found", err)
	}

	replies, filterErr := arn.FilterPosts(func(post *arn.Post) bool {
		post.Text = strings.Replace(post.Text, "http://", "https://", -1)
		return post.ThreadID == thread.ID
	})

	arn.SortPostsLatestLast(replies)

	for i := 0; i < 7; i++ {
		replies = append(replies, replies...)
	}

	println(len(replies))

	// Pre-render markdown
	// flow.Parallel(func() {
	// 	for _, reply := range replies[0:256] {
	// 		reply.HTML()
	// 	}
	// }, func() {
	// 	for _, reply := range replies[256:512] {
	// 		reply.HTML()
	// 	}
	// }, func() {
	// 	for _, reply := range replies[512:768] {
	// 		reply.HTML()
	// 	}
	// }, func() {
	// 	for _, reply := range replies[768:1024] {
	// 		reply.HTML()
	// 	}
	// })

	if filterErr != nil {
		return ctx.Error(500, "Error fetching thread replies", err)
	}

	return ctx.HTML(components.Thread(thread, replies, user))
}
