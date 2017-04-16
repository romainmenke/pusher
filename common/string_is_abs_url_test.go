package common

import "testing"

var absoluteRes = false

func BenchmarkIsAbsoluteA(b *testing.B) {
	for n := 0; n < b.N; n++ {
		absoluteRes = IsAbsolute("/fonts/CutiveMono-Regular.ttf")
	}
}

func BenchmarkIsAbsoluteB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		absoluteRes = IsAbsolute("https://www.foo.com/fonts/CutiveMono-Regular.ttf")
	}
}

func TestIsAbsolute_False(t *testing.T) {
	if IsAbsolute("/fonts/CutiveMono-Regular.ttf") {
		t.Fatal("expected false")
	}
}

func TestIsAbsolute_True_A(t *testing.T) {
	if !IsAbsolute("https://www.foo.com/fonts/CutiveMono-Regular.ttf") {
		t.Fatal("expected true")
	}
}

func TestIsAbsolute_True_B(t *testing.T) {
	if !IsAbsolute("//www.foo.com/fonts/CutiveMono-Regular.ttf") {
		t.Fatal("expected true")
	}
}

func TestIsAbsolute_True_C(t *testing.T) {
	if !IsAbsolute("www.foo.com/fonts/CutiveMono-Regular.ttf") {
		t.Fatal("expected true")
	}
}

func TestIsAbsolute_True_D(t *testing.T) {
	if !IsAbsolute("foo.com/fonts/CutiveMono-Regular.ttf") {
		t.Fatal("expected true")
	}
}

func TestIsAbsolute_True_E(t *testing.T) {
	if !IsAbsolute("fonts/CutiveMono-Regular.ttf") {
		t.Fatal("expected true")
	}
}
