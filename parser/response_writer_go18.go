// +build go1.8

package parser

import (
	"bytes"
	"net/http"
)

type responseWriter struct {
	http.ResponseWriter

	// Code is the HTTP response code set by WriteHeader.
	//
	// Note that if a Handler never calls WriteHeader or Write,
	// this might end up being 0, rather than the implicit
	// http.StatusOK. To get the implicit value, use the Result
	// method.
	statusCode int

	// Body is the buffer to which the Handler's Write calls are sent.
	// If nil, the Writes are silently discarded.
	body *bytes.Buffer

	// Flushed is whether the Handler called Flush.
	flushed bool

	path string
}

func newResponseWriter(w http.ResponseWriter, r *http.Request) *responseWriter {
	return &responseWriter{
		w,
		0,
		new(bytes.Buffer),
		false,
		r.URL.RequestURI(),
	}
}

// Header returns the response headers.
func (w *responseWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

// Write always succeeds and writes to rw.Body, if not nil.
func (w *responseWriter) Write(buf []byte) (int, error) {
	if w.statusCode == 0 {
		w.statusCode = 200
	}

	if w.body != nil {
		l := len(buf)
		if l > 1024 {
			l = 1024
		}
		w.body.Write(buf[:l])
	}

	return w.ResponseWriter.Write(buf)
}

// WriteHeader sets rw.Code. After it is called, changing rw.Header
// will not affect rw.HeaderMap.
func (w *responseWriter) WriteHeader(s int) {
	if w.statusCode == 0 {
		w.statusCode = s
	}

	w.ResponseWriter.WriteHeader(w.statusCode)
}

// Flush sets rw.Flushed to true.
func (w *responseWriter) Flush() {
	flusher, ok := w.ResponseWriter.(http.Flusher)
	if ok && flusher != nil {
		flusher.Flush()
	}
}
