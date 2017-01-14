package link

import (
	"net/http"
	"testing"
)

var testHandler = func(w http.ResponseWriter, r *http.Request) {

	// adding link headers is done manually in the example.
	// this better illustrates the workings of the InitiatePush handler
	switch r.URL.RequestURI() {
	case "/":
		w.Header().Add("Link", "</css/stylesheet.css>; rel=preload; as=style;")
		w.Header().Add("link", "</fonts/CutiveMono-Regular.ttf>; rel=preload; as=font;")
	default:
	}

	w.Write([]byte{})
}

func BenchmarkLinkHandler(b *testing.B) { // 11 allocs

	for n := 0; n < b.N; n++ {

		testReq, testErr := http.NewRequest("GET", "/", nil)
		if testErr != nil {
			b.Fatal(testErr)
		}

		testResponseWriter := newTestWriter()
		testHandlerFunc := HandleFunc(testHandler)

		testHandlerFunc(testResponseWriter, testReq)

	}
}

func BenchmarkRegularHandler(b *testing.B) { // 9 allocs

	for n := 0; n < b.N; n++ {

		testReq, testErr := http.NewRequest("GET", "/", nil)
		if testErr != nil {
			b.Fatal(testErr)
		}

		testResponseWriter := newTestWriter()
		testHandlerFunc := testHandler

		testHandlerFunc(testResponseWriter, testReq)

	}
}

const LinkHeaderKey = "Link"

func BenchmarkAllocA(b *testing.B) { // 3 allocs

	for n := 0; n < b.N; n++ {

		testHeader := http.Header{}
		testHeader[LinkHeaderKey] = []string{"</css/stylesheet.css>; rel=preload; as=style;", "</fonts/CutiveMono-Regular.ttf>; rel=preload; as=font;"}

	}
}

func BenchmarkAllocB(b *testing.B) { // 0 allocs

	for n := 0; n < b.N; n++ {

		testResponseWriter := newTestWriter()
		testResponseWriter.WriteHeader(200)

	}
}

func BenchmarkAllocC(b *testing.B) { // 5 allocs

	for n := 0; n < b.N; n++ {

		testResponseWriter := newTestWriter()
		testHeader := http.Header{}

		testHeader[LinkHeaderKey] = []string{"</css/stylesheet.css>; rel=preload; as=style;", "</fonts/CutiveMono-Regular.ttf>; rel=preload; as=font;", "</blah/foo>"}

		InitiatePush(testHeader, testResponseWriter)

	}
}

func BenchmarkAllocD(b *testing.B) { // 4 allocs

	for n := 0; n < b.N; n++ {

		testResponseWriter := newTestWriter()
		testHeader := http.Header{}

		testHeader[LinkHeaderKey] = []string{}

		InitiatePush(testHeader, testResponseWriter)

	}
}

func BenchmarkAllocE(b *testing.B) { // 5 allocs

	for n := 0; n < b.N; n++ {

		testResponseWriter := newTestWriter()
		testHeader := http.Header{}

		testHeader[LinkHeaderKey] = []string{"</css/stylesheet.css>; rel=preload; as=style;"}

		InitiatePush(testHeader, testResponseWriter)

	}
}

var testGlobalResponseWriter *testWriter
var testGlobalHeader http.Header

func init() {
	testGlobalResponseWriter = newTestWriter()
	testGlobalHeader = http.Header{}
	testGlobalHeader[LinkHeaderKey] = []string{"</css/stylesheet.css>; rel=preload; as=style;", "</fonts/CutiveMono-Regular.ttf>; rel=preload; as=font;", "</blah/foo>"}
}

func BenchmarkAllocF(b *testing.B) { // 0 allocs

	for n := 0; n < b.N; n++ {

		InitiatePush(testGlobalHeader, testGlobalResponseWriter)

	}
}
