// +build go1.8

package casper

import (
	"net/http"

	"github.com/romainmenke/pusher/common"
)

// Handler wraps an http.Handler with H2 Push functionality.
func Handler(p int, n int, handler http.Handler) http.Handler {

	c := &Casper{
		p: uint(p),
		n: uint(n),
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		hashValues, err := c.readCookie(r)
		if err != nil {
			// log.Println(err)
			err = nil
		}

		ctx := r.Context()
		ctx = contextWithCasper(ctx, c)
		r = r.WithContext(ctx)

		// If CanPush returns false, use the input handler.
		// Else -> wrap it.
		if !CanPush(w, r) {
			handler.ServeHTTP(w, r)
			return
		}

		// Get a responseWriter from the sync.Pool.
		rw := getResponseWriter(r.Context(), w, hashValues)
		// defer close the responseWriter.
		// This returns it to the sync.Pool and zeroes all values and pointers.
		defer rw.close()

		handler.ServeHTTP(rw, r)
	})
}

// CanPush checks if the Request is Pushable and the ResponseWriter supports H2 Push.
func CanPush(w http.ResponseWriter, r *http.Request) bool {

	// Only GET requests should trigger Pushes.
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		return false
	}

	// Push is only supported from HTTP2.0.
	if r.ProtoMajor < 2 {
		return false
	}

	// The http.ResponseWriter has to be http.Pusher.
	_, ok := w.(http.Pusher)
	if !ok {
		return false
	}

	// The request must not be proxied.
	// Proxies might not support forwarding Pushes.
	if r.Header.Get(common.XForwardedFor) != "" {
		return false
	}

	return true
}
