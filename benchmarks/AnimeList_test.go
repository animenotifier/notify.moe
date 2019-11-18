package benchmarks

import (
	"testing"

	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
)

func BenchmarkRenderAnimeList(b *testing.B) {
	user, _ := arn.GetUser("4J6qpK1ve")
	animeList := user.AnimeList()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		components.AnimeList(animeList.Items, -1, user, user)
	}
}
