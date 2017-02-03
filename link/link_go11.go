// +build !go1.8

package link

import (
	"log"
	"net/http"
)

// Handler wraps an http.HandlerFunc with H2 Push functionality.
func Handler(handler http.Handler) http.Handler {
	return handler
}

// CanPush checks if the Request is Pushable.
func CanPush(w http.ResponseWriter, r *http.Request) bool {
	return false
}

// InitiatePush parses Link Headers of a response to generate Push Frames.
func InitiatePush(header http.Header, pusher http.Pusher) {
	return
}
