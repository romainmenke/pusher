// +build go1.8

package link

import "net/http"

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

	var splitIndex int
PUSH_LOOP:
	for index, link := range linkHeaders {

		if index > headerAmountLimit {
			break PUSH_LOOP
		}

		pushLink := parseLinkHeader(link)
		if pushLink != "" {

			err := pusher.Push(pushLink, &http.PushOptions{
				Header: w.request.Header,
			})
			if err != nil {
				switch err.Error() {
				case http2ErrPushLimitReached:
					break PUSH_LOOP
				case http2ErrRecursivePush:
					break PUSH_LOOP
				default:
					continue PUSH_LOOP
				}
			}

			linkHeaders[splitIndex], linkHeaders[index] = linkHeaders[index], linkHeaders[splitIndex]
			splitIndex++
		}
	}

	w.ResponseWriter.Header()[Link] = linkHeaders[splitIndex:]
	w.ResponseWriter.Header()[GoH2Pushed] = linkHeaders[:splitIndex]

}
