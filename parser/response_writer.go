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

	headerWritten bool

	request *http.Request

	settings *settings
}

// Header returns the response headers.
func (w *responseWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

// WriteHeader sets rw.Code. After it is called, changing rw.Header
// will not affect rw.HeaderMap.
func (w *responseWriter) WriteHeader(s int) {
	if w.statusCode == 0 && !w.headerWritten {
		w.headerWritten = true
		w.statusCode = s
	}

	w.ResponseWriter.WriteHeader(w.statusCode)
}

func (w *responseWriter) CloseNotify() <-chan bool {
	return w.ResponseWriter.(http.CloseNotifier).CloseNotify()
}

func (w *responseWriter) Flush() {
	flusher, ok := w.ResponseWriter.(http.Flusher)
	if ok && flusher != nil {
		flusher.Flush()
	}
}
