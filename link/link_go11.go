// +build !go1.8

package link

import "net/http"

// Handler wraps an http.HandlerFunc with H2 Push functionality.
func Handler(handler http.Handler, options ...Option) http.Handler {
	return handler
}
