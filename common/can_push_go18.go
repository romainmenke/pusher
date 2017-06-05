package common

import "net/http"

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
	if r.Header.Get(XForwardedFor) != "" {
		return false
	}

	return true
}
