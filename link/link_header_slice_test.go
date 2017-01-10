package link

import "testing"

func TestLinkHeaderSlice(t *testing.T) {
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

	LinkHeaderSlice(header).Sort()

	for _, h := range header {
		t.Log(h)
		t.Log("   --", parseLinkHeader(h))
	}

}

func BenchmarkLinkHeaderSlice(b *testing.B) {

	for n := 0; n < b.N; n++ {

		LinkHeaderSlice(testHeader()).Sort()

	}

}

var testHeader = func() []string {
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
