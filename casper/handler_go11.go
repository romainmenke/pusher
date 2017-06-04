// +build !go1.8

package casper

import "net/http"

func Handler(p int, n int, handler http.Handler) http.Handler {
	return handler
}
