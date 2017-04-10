// +build go1.8

package link

import (
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

// getResponseWriter returns a responseWriter from the sync.Pool.
// as a save guard reset is also called before returning the responseWriter.
func getResponseWriter(w http.ResponseWriter, r *http.Request) *responseWriter {
	rw := writerPool.Get().(*responseWriter)
	rw.reset()

	rw.ResponseWriter = w
	rw.request = r
	return rw
}
