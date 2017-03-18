package link

import "testing"

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

		sortLinkHeaders(testHeaderLink())

	}

}

func BenchmarkLinkHeaderSplit(b *testing.B) {

	for n := 0; n < b.N; n++ {

		splitLinkHeaders(testHeaderLink())

	}

}

var testHeaderLink = func() []string {
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
