package benchmarks

import (
	"testing"

	"github.com/animenotifier/arn"
)

func BenchmarkDBGetMap(b *testing.B) {
	user, _ := arn.GetUser("4J6qpK1ve")

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			animeList, _ := arn.GetAnimeList(user)
			noop(animeList)
		}
	})
}

func BenchmarkDBGet(b *testing.B) {
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
