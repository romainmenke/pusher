package link

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var testRes = []byte{}

var testHandler = func(w http.ResponseWriter, r *http.Request) {

	// adding link headers is done manually in the example.
	// this better illustrates the workings of the push handler
	switch r.URL.RequestURI() {
	case "/":
		w.Header().Add("Link", "</css/stylesheet.css>; rel=preload; as=style;")
		w.Header().Add("link", "</fonts/CutiveMono-Regular.ttf>; rel=preload; as=font;")
	default:
	}

	w.Write(testRes)
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

func BenchmarkPusher(b *testing.B) { // 16 allocs

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

		h(rr, req)

	}

}

func BenchmarkHandler(b *testing.B) { // 15 allocs

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

		h(rr, req)

	}
}

var testReq *http.Request
var testErr error
var testResponseWriter http.ResponseWriter
var testHandlerFunc http.HandlerFunc

func BenchmarkAllocA(b *testing.B) { // 7 allocs

	for n := 0; n < b.N; n++ {

		testReq, testErr = http.NewRequest("GET", "/", nil)
		if testErr != nil {
			b.Fatal(testErr)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		testResponseWriter = httptest.NewRecorder()

		testHandlerFunc = HandlerFunc(testHandler)

	}
}

func BenchmarkAllocB(b *testing.B) { // 6 allocs

	for n := 0; n < b.N; n++ {

		testReq, testErr = http.NewRequest("GET", "/", nil)
		if testErr != nil {
			b.Fatal(testErr)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		testResponseWriter = httptest.NewRecorder()

		testHandlerFunc = testHandler

	}
}

func BenchmarkAllocC(b *testing.B) { // 7 allocs

	for n := 0; n < b.N; n++ {

		testReq, testErr = http.NewRequest("GET", "/", nil)
		if testErr != nil {
			b.Fatal(testErr)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		testResponseWriter = httptest.NewRecorder()

		testHandlerFunc = newPushHandlerFunc(testHandler)

	}
}

func BenchmarkAllocD(b *testing.B) { // 7 allocs

	for n := 0; n < b.N; n++ {

		testReq, testErr = http.NewRequest("GET", "/", nil)
		if testErr != nil {
			b.Fatal(testErr)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		testResponseWriter = httptest.NewRecorder()

		// if the client does not support H2 Push, abort as early as possible
		var ok bool
		_, ok = testResponseWriter.(http.Pusher)
		if !ok {
			return
		}

		var p pusher
		var header http.Header
		header = make(http.Header)
		p = pusher{writer: testResponseWriter, header: header}
		p.Push()

	}
}
