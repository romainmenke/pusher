package link

import "net/http"

// var _ FullResponseWriter = &testRecorder{}

type testWriter struct {
	HeaderMap http.Header // the HTTP response headers

}

func NewTestWriter() *testWriter {
	return &testWriter{
		HeaderMap: make(http.Header),
	}
}

func (rw *testWriter) Header() http.Header {
	m := rw.HeaderMap
	if m == nil {
		m = make(http.Header)
		rw.HeaderMap = m
	}
	return m
}

func (rw *testWriter) writeHeader(b []byte, str string) {

}

func (rw *testWriter) Write(buf []byte) (int, error) {
	return len(buf), nil
}

func (rw *testWriter) WriteHeader(code int) {

}

func (rw *testWriter) Flush() {
}

func (rw *testWriter) Push(path string, options *http.PushOptions) error {
	return nil
}
