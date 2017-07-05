package benchmarks

import (
	"testing"

	"github.com/animenotifier/arn"
)

func BenchmarkDBGetMap(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			arn.DB.GetMap("AnimeList", "4J6qpK1ve")
		}
	})
}

func BenchmarkDBGet(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			arn.DB.Get("AnimeList", "4J6qpK1ve")
		}
	})
}
