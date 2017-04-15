package parser

import "net/http"

// Handler wraps an http.Handler with H2 Push functionality.
func Handler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Get a responseWriter from the sync.Pool.
		var rw = getResponseWriter(w, r)
		// defer close the responseWriter.
		// This returns it to the sync.Pool and zeroes all values and pointers.
		defer rw.close()

		// handle.
		handler.ServeHTTP(rw, r)

	})
}
