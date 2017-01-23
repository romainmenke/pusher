package linkheader

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
)

func TestGetHeaders(t *testing.T) {

	mux := setupTestMux()

	testReq, testErr := http.NewRequest("GET", "/", nil)
	if testErr != nil {
		t.Fatal(testErr)
	}

	fmt.Println(mux.getHeaders(testReq.Host, testReq.URL.Path))

}

func TestRead(t *testing.T) {

	mux := &HeaderMux{
		m:     make(map[string]muxEntry),
		links: make(map[string]struct{}),
		mu:    &sync.RWMutex{},
	}

	err := mux.Read("./example/linkheaders.txt")
	if err != nil {
		t.Fatal(err)
	}

	testReq, testErr := http.NewRequest("GET", "/", nil)
	if testErr != nil {
		t.Fatal(testErr)
	}

	fmt.Println(mux.getHeaders(testReq.Host, testReq.URL.Path))

}

func setupTestMux() *HeaderMux {

	mux := &HeaderMux{
		m:     make(map[string]muxEntry),
		links: make(map[string]struct{}),
		mu:    &sync.RWMutex{},
	}

	mux.Store("/", []string{
		"</css/stylesheet.css>; rel=preload; as=style;",
		"</fonts/CutiveMono-Regular.ttf>; rel=preload; as=font;",
	})
	mux.Store("/alpha.html", []string{
		"</css/stylesheet.css>; rel=preload; as=style;",
		"</fonts/CutiveMono-Regular.ttf>; rel=preload; as=font;",
		"</js/text_change.js>; rel=preload; as=script;",
	})
	mux.Store("/beta.html", []string{
		"</css/stylesheet.css>; rel=preload; as=style;",
		"</fonts/CutiveMono-Regular.ttf>; rel=preload; as=font;",
		"</img/gopher.png>; rel=preload; as=image;",
		"</img/gopher1.png>; rel=preload; as=image;",
		"</img/gopher2.png>; rel=preload; as=image;",
		"</img/gopher3.png>; rel=preload; as=image;",
		"</img/gopher4.png>; rel=preload; as=image;",
		"</img/gopher5.png>; rel=preload; as=image;",
	})
	mux.Store("/gamma.html", []string{
		"</css/stylesheet.css>; rel=preload; as=style;",
		"</fonts/CutiveMono-Regular.ttf>; rel=preload; as=font;",
		"</call.json>; rel=preload;",
	})

	return mux

}
