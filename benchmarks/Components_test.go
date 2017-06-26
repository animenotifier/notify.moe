package benchmarks

import (
	"testing"

	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

func BenchmarkThread(b *testing.B) {
	thread, _ := arn.GetThread("HJgS7c2K")
	thread.HTML() // Pre-render markdown

	replies, _ := arn.FilterPosts(func(post *arn.Post) bool {
		post.HTML() // Pre-render markdown
		return post.ThreadID == thread.ID
	})

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			components.Thread(thread, replies, nil)
		}
	})
}
