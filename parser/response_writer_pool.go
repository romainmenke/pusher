package parser

import (
	"bytes"
	"net/http"
	"sync"
)

// writerPool is the sync.Pool used to reduce GC pauses
var writerPool *sync.Pool

// init initialises the writerPool
func init() {
	writerPool = &sync.Pool{
		New: func() interface{} {
			return &responseWriter{}
		},
	}
}

func (w *responseWriter) reset() {
	w.body = nil
	w.headerWritten = false
	w.request = nil
	w.ResponseWriter = nil
	w.settings = nil
	w.statusCode = 0
}

func (w *responseWriter) close() {
	w.reset()
	writerPool.Put(w)
}

// getResponseWriter returns a responseWriter from the sync.Pool.
// as a save guard reset is also called before returning the responseWriter.
func getResponseWriter(s *settings, w http.ResponseWriter, r *http.Request) *responseWriter {
	rw := writerPool.Get().(*responseWriter)
	rw.reset()

	rw.body = new(bytes.Buffer)
	rw.request = r
	rw.ResponseWriter = w
	rw.settings = s
	return rw
}
