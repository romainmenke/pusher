package link

import (
	"net/http"
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

var testReq *http.Request
var testErr error
var testResponseWriter http.ResponseWriter
var testHandlerFunc http.HandlerFunc

func TestPusher(t *testing.T) {

	testReq, testErr := http.NewRequest("GET", "/", nil)
	if testErr != nil {
		t.Fatal(testErr)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	testResponseWriter := NewTestWriter()
	handler := HandlerFunc(testHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(testResponseWriter, testReq)

}

func BenchmarkPusher(b *testing.B) { // 12 allocs

	for n := 0; n < b.N; n++ {

		testReq, testErr = http.NewRequest("GET", "/", nil)
		if testErr != nil {
			b.Fatal(testErr)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		testResponseWriter = NewTestWriter()

		testHandlerFunc = HandlerFunc(testHandler)

		testHandlerFunc(testResponseWriter, testReq)

	}

}

func BenchmarkRegularHandler(b *testing.B) { // 9 allocs

	for n := 0; n < b.N; n++ {

		testReq, testErr = http.NewRequest("GET", "/", nil)
		if testErr != nil {
			b.Fatal(testErr)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		testResponseWriter = NewTestWriter()

		testHandlerFunc = testHandler

		testHandlerFunc(testResponseWriter, testReq)

	}
}

func BenchmarkAllocA(b *testing.B) { // 6 allocs

	for n := 0; n < b.N; n++ {

		testReq, testErr = http.NewRequest("GET", "/", nil)
		if testErr != nil {
			b.Fatal(testErr)
		}

		testResponseWriter = NewTestWriter()

		testHandlerFunc = HandlerFunc(testHandler)

	}
}

func BenchmarkAllocB(b *testing.B) { // 5 allocs

	for n := 0; n < b.N; n++ {

		testReq, testErr = http.NewRequest("GET", "/", nil)
		if testErr != nil {
			b.Fatal(testErr)
		}

		testResponseWriter = NewTestWriter()

		testHandlerFunc = testHandler

	}
}

func BenchmarkAllocC(b *testing.B) { // 6 allocs

	for n := 0; n < b.N; n++ {

		testReq, testErr = http.NewRequest("GET", "/", nil)
		if testErr != nil {
			b.Fatal(testErr)
		}

		testResponseWriter = NewTestWriter()

		testHandlerFunc = newPushHandlerFunc(testHandler)

	}
}

func BenchmarkAllocD(b *testing.B) { // 10 allocs

	for n := 0; n < b.N; n++ {

		testReq, testErr = http.NewRequest("GET", "/", nil)
		if testErr != nil {
			b.Fatal(testErr)
		}

		testResponseWriter = NewTestWriter()

		var ok bool
		_, ok = testResponseWriter.(http.Pusher)
		if !ok {
			b.Fatal("no pusher")
		}

		var p pusher
		var header http.Header
		header = make(http.Header)
		p = pusher{writer: testResponseWriter, header: header}

		p.Header()["Link"] = []string{"</css/stylesheet.css>; rel=preload; as=style;", "</fonts/CutiveMono-Regular.ttf>; rel=preload; as=font;"}

		p.Push()

	}
}
