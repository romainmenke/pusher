package link

import (
	"fmt"
	"testing"
)

func TestParseLinkHeaderA(t *testing.T) {

	res := parseLinkHeader("</fonts/CutiveMono-Regular.ttf>; rel=preload; as=font;")
	if res != "/fonts/CutiveMono-Regular.ttf" {
		t.Fatal("expected : /fonts/CutiveMono-Regular.ttf , got :", res)
	}
}

func TestParseLinkHeaderB(t *testing.T) {

	res := parseLinkHeader("</fonts/CutiveMono-Regular.ttf>; as=font;")
	if res != "" {
		t.Fatal("expected : <empty string> , got :", res)
	}
}

func TestParseLinkHeaderC(t *testing.T) {

	res := parseLinkHeader("</fonts/CutiveMono-Regular.ttf>; rel=preload; as=font; nopush;")
	if res != "" {
		t.Fatal("expected : <empty string> , got :", res)
	}
}

func TestParseLinkHeaderE(t *testing.T) {

	res := parseLinkHeader("< /fonts/CutiveMono-Regular.ttf>; rel=preload; as=font;")
	if res != "/fonts/CutiveMono-Regular.ttf" {
		t.Fatal("expected : /fonts/CutiveMono-Regular.ttf , got :", res)
	}
}

func TestParseLinkHeaderF(t *testing.T) {

	res := parseLinkHeader("</fonts/CutiveMono-Regular.ttf >; rel=preload; as=font;")
	if res != "/fonts/CutiveMono-Regular.ttf" {
		t.Fatal("expected : /fonts/CutiveMono-Regular.ttf , got :", res)
	}
}

func TestParseLinkHeaderBadA(t *testing.T) {

	res := parseLinkHeader("/fonts/CutiveMono-Regular.ttf>; rel=preload; as=font;")
	if res != "" {
		t.Fatal("expected : <empty string> , got :", res)
	}
}

func TestParseLinkHeaderBadB(t *testing.T) {

	res := parseLinkHeader("</fonts/CutiveMono-Regular.ttf; rel=preload; as=font;")
	if res != "" {
		t.Fatal("expected : <empty string> , got :", res)
	}
}

func TestParseLinkHeaderBadC(t *testing.T) {

	res := parseLinkHeader("</aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.ttf; rel=preload; as=font;")
	if res != "" {
		t.Fatal("expected : <empty string> , got :", res)
	}
}

var parseLinkHeaderRes = ""

func BenchmarkParseLinkHeader_1(b *testing.B) {
	parseLinkHeaderBenchFactory(0)(b)
}

func BenchmarkParseLinkHeader_10(b *testing.B) {
	parseLinkHeaderBenchFactory(10)(b)
}

func BenchmarkParseLinkHeader_100(b *testing.B) {
	parseLinkHeaderBenchFactory(100)(b)
}

func BenchmarkParseLinkHeader_1000(b *testing.B) {
	parseLinkHeaderBenchFactory(1000)(b)
}

func BenchmarkParseLinkHeader_10000(b *testing.B) {
	parseLinkHeaderBenchFactory(10000)(b)
}

func parseLinkHeaderBenchFactory(length int) func(b *testing.B) {
	return func(b *testing.B) {
		testString := ""
		for i := 0; i < length; i++ {
			testString += "a"
		}
		testString = fmt.Sprintf("</%s>; rel=preload;", testString)

		b.ResetTimer()

		for n := 0; n < b.N; n++ {
			parseLinkHeaderRes = parseLinkHeader(testString)
		}
	}
}
