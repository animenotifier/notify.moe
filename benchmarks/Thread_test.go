package benchmarks

import (
	"testing"

	"github.com/animenotifier/notify.moe/arn"
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

	for i := 0; i < b.N; i++ {
		components.Thread(thread, nil)
	}
}
