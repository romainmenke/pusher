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
		absoluteRes = isAbsolute("https://www.foo.com/fonts/CutiveMono-Regular.ttf")
	}
}

func TestIsAbsolute_False(t *testing.T) {
	if isAbsolute("/fonts/CutiveMono-Regular.ttf") {
		t.Fatal("expected false")
	}
}

func TestIsAbsolute_True_A(t *testing.T) {
	if !isAbsolute("https://www.foo.com/fonts/CutiveMono-Regular.ttf") {
		t.Fatal("expected true")
	}
}

func TestIsAbsolute_True_B(t *testing.T) {
	if !isAbsolute("//www.foo.com/fonts/CutiveMono-Regular.ttf") {
		t.Fatal("expected true")
	}
}

func TestIsAbsolute_True_C(t *testing.T) {
	if !isAbsolute("www.foo.com/fonts/CutiveMono-Regular.ttf") {
		t.Fatal("expected true")
	}
}

func TestIsAbsolute_True_D(t *testing.T) {
	if !isAbsolute("foo.com/fonts/CutiveMono-Regular.ttf") {
		t.Fatal("expected true")
	}
}

func TestIsAbsolute_True_E(t *testing.T) {
	if !isAbsolute("fonts/CutiveMono-Regular.ttf") {
		t.Fatal("expected true")
	}
}
