package pusher

import "net/http"

// Pusher wraps a http.HandlerFunc.
// It will automatically generate Push Promises
func Pusher(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer handler(w, r)

		setInitiatorForWriter(w, r)

		if !isPageGet(r) {
			go addToPushMap(r)
			return
		}

		if p, ok := w.(http.Pusher); ok {
			readFromPushMap(r.RequestURI, func(path string) {
				opts := &http.PushOptions{}
				opts = setInitiatorForOptions(r, opts)
				p.Push(path, opts)
			})
		}

	})
}
