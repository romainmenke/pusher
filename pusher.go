package pusher

import "net/http"

// Handler wraps a http.HandlerFunc.
// It will automatically generate Push Promises
func Handler(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer handler(w, r)

		setInitiatorForWriter(w, r)
		if !isPageGet(r) {
			if getInitiator(r) != "" {
				go addToPushMap(r, 1)
			} else {
				go addToPushMap(r, 0.1)
			}
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
