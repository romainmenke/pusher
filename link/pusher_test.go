package link

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var testHandler = func(w http.ResponseWriter, r *http.Request) {

	// adding link headers is done manually in the example.
	// this better illustrates the workings of the push handler
	switch r.URL.RequestURI() {
	case "/":
		w.Header().Add("Link", "</css/stylesheet.css>; rel=preload; as=style;")
		w.Header().Add("link", "</fonts/CutiveMono-Regular.ttf>; rel=preload; as=font;")
	default:
	}

	w.Write([]byte{})
}

func TestPusher(t *testing.T) {

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := HandlerFunc(testHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

}

func BenchmarkPusher(b *testing.B) {

	for n := 0; n < b.N; n++ {

		var (
			req *http.Request
			err error
			rr  http.ResponseWriter
			h   http.HandlerFunc
		)

		req, err = http.NewRequest("GET", "/", nil)
		if err != nil {
			b.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr = httptest.NewRecorder()

		h = HandlerFunc(testHandler)

		h(rr, req) // 16 allocs

	}

}

func BenchmarkHandler(b *testing.B) {

	for n := 0; n < b.N; n++ {

		var (
			req *http.Request
			err error
			rr  http.ResponseWriter
			h   http.HandlerFunc
		)

		req, err = http.NewRequest("GET", "/", nil)
		if err != nil {
			b.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr = httptest.NewRecorder()

		h = testHandler

		h(rr, req) // 15 allocs

	}
}
