// +build !go1.8

package casper

import "net/http"

func Handler(handler http.Handler) http.Handler {
	return handler
}
