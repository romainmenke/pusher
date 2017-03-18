package link

import "testing"

func BenchmarkIsAbsoluteA(b *testing.B) {

	res := false

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			res = isAbsolute("/fonts/CutiveMono-Regular.ttf")
		}
	})
}

func BenchmarkIsAbsoluteB(b *testing.B) {

	res := false

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			res = isAbsolute("https://www.foo,com/fonts/CutiveMono-Regular.ttf")
		}
	})
}
