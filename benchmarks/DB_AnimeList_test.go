package benchmarks

import (
	"testing"

	"github.com/animenotifier/notify.moe/arn"
)

func BenchmarkDatabaseGetAnimeList(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			list, _ := arn.DB.Get("AnimeList", "4J6qpK1ve")
			animeList := list.(*arn.AnimeList)
			noop(animeList)
		}
	})
}

func noop(list *arn.AnimeList) {}
