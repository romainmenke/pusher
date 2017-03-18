package link

import "testing"

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

func BenchmarkParseLinkHeader(b *testing.B) {

	res := ""

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			res = parseLinkHeader("</fonts/CutiveMono-Regular.ttf>; rel=preload; as=font;")
		}
	})
}
