// +build go1.8

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
	w.ResponseWriter = nil
	w.request = nil
	w.statusCode = 0
	w.extracted = false
}

func (w *responseWriter) close() {
	w.reset()
	writerPool.Put(w)
}

func newResponseWriter(w http.ResponseWriter, r *http.Request) *responseWriter {
	return &responseWriter{
		w,
		0,
		new(bytes.Buffer),
		false,
		r,
	}
}

// getResponseWriter returns a responseWriter from the sync.Pool.
// as a save guard reset is also called before returning the responseWriter.
func getResponseWriter(w http.ResponseWriter, r *http.Request) *responseWriter {
	rw := writerPool.Get().(*responseWriter)
	rw.reset()

	rw.body = new(bytes.Buffer)
	rw.ResponseWriter = w
	rw.request = r
	return rw
}
