package benchmarks

import (
	"testing"

	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

func BenchmarkRenderThread(b *testing.B) {
	thread, _ := arn.GetThread("HJgS7c2K")
	thread.HTML() // Pre-render markdown
	replies := thread.Posts()

	for _, reply := range replies {
		reply.HTML() // Pre-render markdown
	}

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			components.Thread(thread, nil)
		}
	})
}

func BenchmarkRenderAnimeList(b *testing.B) {
	user, _ := arn.GetUser("4J6qpK1ve")
	animeList := user.AnimeList()

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			components.AnimeList(animeList.Items, -1, user, user)
		}
	})
}
