// +build go1.8

package casper

import (
	"net/http"

	"github.com/romainmenke/pusher/common"
)

// Handler wraps an http.Handler with H2 Push functionality.
func Handler(handler http.Handler, options ...Option) http.Handler {

	c := &Casper{
		p:        uint(common.PushAmountLimit * common.PushAmountLimit),
		n:        uint(common.PushAmountLimit),
		settings: settings{},
	}

	for _, opt := range options {
		opt(&c.settings)
	}

	if c.settings.cookieMaxAge == 0 {
		c.settings.cookieMaxAge = 3600
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		ctx = contextWithCasper(ctx, c)
		r = r.WithContext(ctx)

		// If CanPush returns false, use the input handler.
		// Else -> wrap it.
		if !common.CanPush(w, r) {
			handler.ServeHTTP(w, r)
			return
		}

		// Get a responseWriter from the sync.Pool.
		rw := getResponseWriter(r.Context(), w)
		// defer close the responseWriter.
		// This returns it to the sync.Pool and zeroes all values and pointers.
		defer rw.close()

		err := c.readCookie(r, &rw.hashValues)
		if err != nil {
			// log.Println(err)
			err = nil
		}

		handler.ServeHTTP(rw, r)
	})
}
