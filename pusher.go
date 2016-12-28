package pusher

import "net/http"

// Pusher wraps a http.HandlerFunc.
// It will automatically generate Push Promises
func Pusher(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer handler(w, r)

		if !isPageGet(r) {
			go addToPushMap(r.Referer(), r.URL.String())

			return
		}

		if p, ok := w.(http.Pusher); ok {
			readFromPushMap(r.URL.String(), func(path string) {
				p.Push(path, nil)
			})
		}

	})
}
