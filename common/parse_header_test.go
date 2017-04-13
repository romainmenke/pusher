package common

import (
	"fmt"
	"testing"
)

func TestParseLinkHeaderA(t *testing.T) {

	res := ParseLinkHeader("</fonts/CutiveMono-Regular.ttf>; rel=preload; as=font;")
	if res != "/fonts/CutiveMono-Regular.ttf" {
		t.Fatal("expected : /fonts/CutiveMono-Regular.ttf , got :", res)
	}
}

func TestParseLinkHeaderB(t *testing.T) {

	res := ParseLinkHeader("</fonts/CutiveMono-Regular.ttf>; as=font;")
	if res != "" {
		t.Fatal("expected : <empty string> , got :", res)
	}
}

func TestParseLinkHeaderC(t *testing.T) {

	res := ParseLinkHeader("</fonts/CutiveMono-Regular.ttf>; rel=preload; as=font; nopush;")
	if res != "" {
		t.Fatal("expected : <empty string> , got :", res)
	}
}

func TestParseLinkHeaderE(t *testing.T) {

	res := ParseLinkHeader("< /fonts/CutiveMono-Regular.ttf>; rel=preload; as=font;")
	if res != "/fonts/CutiveMono-Regular.ttf" {
		t.Fatal("expected : /fonts/CutiveMono-Regular.ttf , got :", res)
	}
}

func TestParseLinkHeaderF(t *testing.T) {

	res := ParseLinkHeader("</fonts/CutiveMono-Regular.ttf >; rel=preload; as=font;")
	if res != "/fonts/CutiveMono-Regular.ttf" {
		t.Fatal("expected : /fonts/CutiveMono-Regular.ttf , got :", res)
	}
}

func TestParseLinkHeaderBadA(t *testing.T) {

	res := ParseLinkHeader("/fonts/CutiveMono-Regular.ttf>; rel=preload; as=font;")
	if res != "" {
		t.Fatal("expected : <empty string> , got :", res)
	}
}

func TestParseLinkHeaderBadB(t *testing.T) {

	res := ParseLinkHeader("</fonts/CutiveMono-Regular.ttf; rel=preload; as=font;")
	if res != "" {
		t.Fatal("expected : <empty string> , got :", res)
	}
}

func TestParseLinkHeaderBadC(t *testing.T) {

	res := ParseLinkHeader("<foo.com/fonts/CutiveMono-Regular.ttf>; rel=preload; as=font;")
	if res != "" {
		t.Fatal("expected : <empty string> , got :", res)
	}
}

func TestParseLinkHeaderLength(t *testing.T) {

	h1025 := "</aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.ttf>; rel=preload; as=font;"

	if len(h1025) != 1025 {
		t.Fatal(len(h1025))
	}

	res := ParseLinkHeader(h1025)
	if res != "" {
		t.Fatal("expected : <empty string> , got :", res)
	}

	h1024 := "</aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.ttf>; rel=preload; as=font;"

	h1024Res := "/aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.ttf"

	if len(h1024) != 1024 {
		t.Fatal(len(h1024))
	}

	res = ParseLinkHeader(h1024)
	if res != h1024Res {
		t.Fatal("expected : url , got :", res)
	}
}

var ParseLinkHeaderRes = ""

func BenchmarkParseLinkHeader_1(b *testing.B) {
	ParseLinkHeaderBenchFactory(0)(b)
}

func BenchmarkParseLinkHeader_10(b *testing.B) {
	ParseLinkHeaderBenchFactory(10)(b)
}

func BenchmarkParseLinkHeader_100(b *testing.B) {
	ParseLinkHeaderBenchFactory(100)(b)
}

func BenchmarkParseLinkHeader_1000(b *testing.B) {
	ParseLinkHeaderBenchFactory(1000)(b)
}

func BenchmarkParseLinkHeader_10000(b *testing.B) {
	ParseLinkHeaderBenchFactory(10000)(b)
}

func ParseLinkHeaderBenchFactory(length int) func(b *testing.B) {
	return func(b *testing.B) {
		testString := ""
		for i := 0; i < length; i++ {
			testString += "a"
		}
		testString = fmt.Sprintf("</%s>; rel=preload;", testString)

		b.ResetTimer()

		for n := 0; n < b.N; n++ {
			ParseLinkHeaderRes = ParseLinkHeader(testString)
		}
	}
}
