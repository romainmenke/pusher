// +build go1.8

package link

import (
	"net/http"

	"github.com/romainmenke/pusher/common"
)

// Handler wraps an http.Handler with H2 Push functionality.
func Handler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// If CanPush returns false, use the input handler.
		// Else -> wrap it.
		if !common.CanPush(w, r) {
			handler.ServeHTTP(w, r)
			return
		}

		// Get a responseWriter from the sync.Pool.
		rw := getResponseWriter(w, r)
		// defer close the responseWriter.
		// This returns it to the sync.Pool and zeroes all values and pointers.
		defer rw.close()

		handler.ServeHTTP(rw, r)
	})
}

// InitiatePush parses Link Headers of a response to generate Push Frames.
func InitiatePush(w *responseWriter) { // 0 allocs

	// Nil checks, these might be redundant.
	if w == nil || w.request == nil {
		return
	}

	// Get the Link Header values from the Response Header.
	linkHeaders, ok := w.Header()[common.Link]
	if !ok {
		return
	}

	pushHeader := http.Header{}
	copyPushSafeHeader(pushHeader, w.request.Header)

	// splitIndex is used to separate Link and Push values without creating a new []string{}.
	var splitIndex int

PUSH_LOOP:
	for index, link := range linkHeaders {

		// Limit the number of values parsed.
		// This is not based on how many are eventually pushed.
		if index >= common.PushAmountLimit {
			break PUSH_LOOP
		}

		// Parse the Link Header Value.
		// This will return either an empty string or a relative url.
		// When not empty -> Push.
		pushLink := common.ParseLinkHeader(link)
		if pushLink != "" {

			// Attempt to send a Push.
			// Pass the original Request Headers by reference.
			err := w.Push(pushLink, &http.PushOptions{
				Header: pushHeader,
				Method: w.request.Method,
			})

			// Handle Push err.
			if err != nil {
				switch err.Error() {

				// No more pushes can be send.
				case http2ErrPushLimitReached:
					break PUSH_LOOP

				// No more pushes can be send.
				case http2ErrRecursivePush:
					break PUSH_LOOP

				// Something went wrong, but maybe nothing serious. Try another Push.
				default:
					continue PUSH_LOOP
				}
			}

			// Swap two header values.
			// This will group all Pushed values to the front of the slice.
			linkHeaders[splitIndex], linkHeaders[index] = linkHeaders[index], linkHeaders[splitIndex]
			splitIndex++
		}
	}

	// Move the pushed values to a new Key to prevent the browser from requesting it.
	w.ResponseWriter.Header()[common.GoH2Pushed] = linkHeaders[:splitIndex]
	// Update 'Link' with the remaining values.
	w.ResponseWriter.Header()[common.Link] = linkHeaders[splitIndex:]

}
