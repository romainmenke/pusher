package link

import "testing"

var absoluteRes = false

func BenchmarkIsAbsoluteA(b *testing.B) {
	for n := 0; n < b.N; n++ {
		absoluteRes = isAbsolute("/fonts/CutiveMono-Regular.ttf")
	}
}

func BenchmarkIsAbsoluteB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		absoluteRes = isAbsolute("https://www.foo,com/fonts/CutiveMono-Regular.ttf")
	}
}
