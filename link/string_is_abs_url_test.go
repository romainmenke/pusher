package link

import "testing"

func BenchmarkIsAbsolute(b *testing.B) {

	res := false

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			res = isAbsolute("/fonts/CutiveMono-Regular.ttf")
		}
	})
}
