// +build go1.8

package link

import "net/http"

const (
	GoH2Pushed    = "Go-H2-Pushed"
	XForwardedFor = "X-Forwarded-For"
	Link          = "Link"
	Get           = "GET"
)

var (
	http2ErrRecursivePush    = "http2: recursive push not allowed"
	http2ErrPushLimitReached = "http2: push would exceed peer's SETTINGS_MAX_CONCURRENT_STREAMS"
)

// Handler wraps an http.Handler with H2 Push functionality.
func Handler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if !CanPush(w, r) {
			handler.ServeHTTP(w, r)
			return
		}

		var rw = getResponseWriter(w, r)
		defer rw.close()

		handler.ServeHTTP(rw, r)

	})
}

// CanPush checks if the Request is Pushable and the ResponseWriter supports H2 Push.
func CanPush(w http.ResponseWriter, r *http.Request) bool {

	if r.Method != Get {
		return false
	}

	if r.ProtoMajor < 2 {
		return false
	}

	_, ok := w.(http.Pusher)
	if !ok {
		return false
	}

	if r.Header.Get(XForwardedFor) != "" {
		return false
	}

	return true
}

// InitiatePush parses Link Headers of a response to generate Push Frames.
func InitiatePush(w *responseWriter) { // 0 allocs

	if w == nil || w.request == nil {
		return
	}

	pusher, ok := w.ResponseWriter.(http.Pusher)
	if !ok {
		return
	}

	linkHeaders, ok := w.Header()[Link]
	if !ok {
		return
	}

	toPush, toLink := splitLinkHeadersAndParse(linkHeaders)

PUSH_LOOP:
	for index, link := range toPush {
		if link == "" {
			continue
		}

		err := pusher.Push(link, &http.PushOptions{
			Header: w.request.Header,
		})
		if err != nil {

			toPush = append(toPush[:index], toPush[index+1:]...)

			switch err.Error() {
			case http2ErrPushLimitReached:
				break PUSH_LOOP
			case http2ErrRecursivePush:
				break PUSH_LOOP
			}
		}
	}

	// leave this in for now
	_ = toLink
	// w.ResponseWriter.Header()[Link] = toLink

	w.ResponseWriter.Header()[GoH2Pushed] = toPush

}
