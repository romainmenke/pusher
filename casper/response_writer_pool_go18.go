// +build go1.8

package casper

import (
	"context"
	"net/http"
	"sync"

	"github.com/romainmenke/pusher/common"
)

// writerPool is the sync.Pool used to reduce GC pauses
var writerPool *sync.Pool

// init initialises the writerPool
func init() {
	writerPool = &sync.Pool{
		New: func() interface{} {
			return &responseWriter{
				hashValues: make([]uint, 0, common.PushAmountLimit),
			}
		},
	}
}

// getResponseWriter returns a responseWriter from the sync.Pool.
// as a save guard reset is also called before returning the responseWriter.
func getResponseWriter(ctx context.Context, w http.ResponseWriter) *responseWriter {
	rw := writerPool.Get().(*responseWriter)
	rw.reset()

	rw.ResponseWriter = w
	rw.ctx = ctx
	return rw
}
