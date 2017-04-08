package link

import (
	"fmt"
	"testing"
)

func TestByPushable(t *testing.T) {
	// https://tools.ietf.org/html/rfc5988
	header := []string{
		"<http://example.com/TheBook/chapter2>; rel=previous; title=previous chapter",
		"</>; rel=http://example.net/foo",
		"</TheBook/chapter2>; rel=previous; title*=UTF-8'de'letztes%20Kapitel",
		"</TheBook/chapter4>; rel=next; title*=UTF-8'de'n%c3%a4chstes%20Kapitel",
		"<http://example.org/>; rel=start http://example.net/relation/other",
		"</fonts/CutiveMono-Regular.ttf>; rel=preload; as=font;",
		"</css/stylesheet.css>; rel=preload; as=style;",
		"</js/text_change.js>; rel=preload; as=script;",
		"</img/gopher.png>; rel=preload; as=image;",
		"</img/gopher2.png>; rel=preload; as=image; nopush;",
		"</call.json>; rel=preload;",
	}

	sortLinkHeaders(header)

	for index, h := range header {
		switch index {
		case 0:
			if h != "</fonts/CutiveMono-Regular.ttf>; rel=preload; as=font;" {
				t.Fail()
			}
			if parseLinkHeader(h) != "/fonts/CutiveMono-Regular.ttf" {
				t.Fail()
			}
		case 1:
			if h != "</css/stylesheet.css>; rel=preload; as=style;" {
				t.Fail()
			}
			if parseLinkHeader(h) != "/css/stylesheet.css" {
				t.Fail()
			}
		case 2:
			if h != "</js/text_change.js>; rel=preload; as=script;" {
				t.Fail()
			}
			if parseLinkHeader(h) != "/js/text_change.js" {
				t.Fail()
			}
		case 3:
			if h != "</img/gopher.png>; rel=preload; as=image;" {
				t.Fail()
			}
			if parseLinkHeader(h) != "/img/gopher.png" {
				t.Fail()
			}
		case 4:
			if h != "</call.json>; rel=preload;" {
				t.Fail()
			}
			if parseLinkHeader(h) != "/call.json" {
				t.Fail()
			}
		case 5:
			if h != "<http://example.com/TheBook/chapter2>; rel=previous; title=previous chapter" {
				t.Fail()
			}
			if parseLinkHeader(h) != "" {
				t.Fail()
			}
		case 6:
			if h != "</>; rel=http://example.net/foo" {
				t.Fail()
			}
			if parseLinkHeader(h) != "" {
				t.Fail()
			}
		case 7:
			if h != "</TheBook/chapter2>; rel=previous; title*=UTF-8'de'letztes%20Kapitel" {
				t.Fail()
			}
			if parseLinkHeader(h) != "" {
				t.Fail()
			}
		case 8:
			if h != "</TheBook/chapter4>; rel=next; title*=UTF-8'de'n%c3%a4chstes%20Kapitel" {
				t.Fail()
			}
			if parseLinkHeader(h) != "" {
				t.Fail()
			}
		case 9:
			if h != "</img/gopher2.png>; rel=preload; as=image; nopush;" {
				t.Fail()
			}
			if parseLinkHeader(h) != "" {
				t.Fail()
			}
		case 10:
			if h != "<http://example.org/>; rel=start http://example.net/relation/other" {
				t.Fail()
			}
			if parseLinkHeader(h) != "" {
				t.Fail()
			}
		}
	}
}

func BenchmarkByPushableSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sortLinkHeaders(testHeaderLinkA())
	}
}

func BenchmarkLinkHeaderSplit(b *testing.B) {
	for n := 0; n < b.N; n++ {
		splitLinkHeaders(testHeaderLinkA())
	}
}

func BenchmarkSplitLinkHeadersAndParse_10_Links(b *testing.B) {
	data := testHeaderLinkNumberScale(10)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {

		tmp := make([]string, len(data))
		copy(tmp, data)

		splitLinkHeadersAndParse(tmp)
	}
}

func BenchmarkSplitLinkHeadersAndParse_100_Links(b *testing.B) {
	data := testHeaderLinkNumberScale(100)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {

		tmp := make([]string, len(data))
		copy(tmp, data)

		splitLinkHeadersAndParse(tmp)
	}
}

func BenchmarkSplitLinkHeadersAndParse_1000_Links(b *testing.B) {
	data := testHeaderLinkNumberScale(1000)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {

		tmp := make([]string, len(data))
		copy(tmp, data)

		splitLinkHeadersAndParse(tmp)
	}
}

func BenchmarkSplitLinkHeadersAndParse_10000_Links(b *testing.B) {
	data := testHeaderLinkNumberScale(10000)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {

		tmp := make([]string, len(data))
		copy(tmp, data)

		splitLinkHeadersAndParse(tmp)
	}
}

func BenchmarkSplitLinkHeadersAndParse_10000_Links_Baseline(b *testing.B) {
	data := testHeaderLinkNumberScale(10000)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {

		tmp := make([]string, len(data))
		copy(tmp, data)
	}
}

func BenchmarkSplitLinkHeadersAndParse_10_Chars(b *testing.B) {
	data := testHeaderLinkLengthScale(10)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {

		tmp := make([]string, len(data))
		copy(tmp, data)

		splitLinkHeadersAndParse(tmp)
	}
}

func BenchmarkSplitLinkHeadersAndParse_100__Chars(b *testing.B) {
	data := testHeaderLinkLengthScale(100)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {

		tmp := make([]string, len(data))
		copy(tmp, data)

		splitLinkHeadersAndParse(tmp)
	}
}

func BenchmarkSplitLinkHeadersAndParse_1000__Chars(b *testing.B) {
	data := testHeaderLinkLengthScale(1000)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {

		tmp := make([]string, len(data))
		copy(tmp, data)

		splitLinkHeadersAndParse(tmp)
	}
}

func BenchmarkSplitLinkHeadersAndParse_10000__Chars(b *testing.B) {
	data := testHeaderLinkLengthScale(10000)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {

		tmp := make([]string, len(data))
		copy(tmp, data)

		splitLinkHeadersAndParse(tmp)
	}
}

func BenchmarkSplitLinkHeadersAndParse_10000__Chars_Baseline(b *testing.B) {
	data := testHeaderLinkLengthScale(10000)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {

		tmp := make([]string, len(data))
		copy(tmp, data)
	}
}

var testHeaderLinkA = func() []string {
	return []string{
		"<http://example.com/TheBook/chapter2>; rel=previous; title=previous chapter",
		"</>; rel=http://example.net/foo",
		"</TheBook/chapter2>; rel=previous; title*=UTF-8'de'letztes%20Kapitel",
		"</TheBook/chapter4>; rel=next; title*=UTF-8'de'n%c3%a4chstes%20Kapitel",
		"<http://example.org/>; rel=start http://example.net/relation/other",
		"</fonts/CutiveMono-Regular.ttf>; rel=preload; as=font;",
		"</css/stylesheet.css>; rel=preload; as=style;",
		"</js/text_change.js>; rel=preload; as=script;",
		"</img/gopher.png>; rel=preload; as=image;",
		"</img/gopher2.png>; rel=preload; as=image; nopush;",
		"</call.json>; rel=preload;",
	}
}

var testHeaderLinkNumberScale = func(amount int) []string {
	headers := make([]string, amount)
	for index := 0; index < amount; index++ {
		headers[index] = fmt.Sprintf("</img/gopher%d.png>; rel=preload; as=image;", index)
	}
	return headers
}

var testHeaderLinkLengthScale = func(length int) []string {
	headers := make([]string, 10)
	for index := 0; index < 10; index++ {
		var part = ""
		for l := 0; l < length; l++ {
			part += "z"
		}
		headers[index] = fmt.Sprintf("</img/gopher-%s.png>; rel=preload; as=image;", part)
	}
	return headers
}
