// +build go1.8

package link

import (
	"net/http"
	"sync"
)

var writerPool *sync.Pool

func init() {
	writerPool = &sync.Pool{
		New: func() interface{} {
			return &responseWriter{}
		},
	}
}

func getResponseWriter(w http.ResponseWriter, r *http.Request) *responseWriter {
	rw := writerPool.Get().(*responseWriter)
	rw.reset()

	rw.ResponseWriter = w
	rw.request = r
	return rw
}
