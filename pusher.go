package pusher

import "net/http"

type pushHandler struct {
	handlerFunc http.HandlerFunc
}

func (h *pushHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handlerFunc(w, r)
}

// Handler wraps a http.Handler.
// It will automatically generate Push Promises
func Handler(handler http.Handler) http.Handler {
	return &pushHandler{
		handlerFunc: newHandlerFunc(handler.ServeHTTP),
	}
}

// HandlerFunc wraps a http.HandlerFunc.
// It will automatically generate Push Promises
func HandlerFunc(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return newHandlerFunc(handlerFunc)
}

func newHandlerFunc(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Note 1 :
		// r/http.Request contains information about the origin of the requested resource
		//
		// Note 2 :
		// w.Header() will contain information about the current requested resource after handling
		//
		// r/http.Request should be used to abort push behavior early

		defer handler(w, r)

		setInitiatorForWriter(w, r)
		if isAssetGet(r) {
			if getPushInitiator(r) != "" { // pushed content
				addToPushMap(r, 0.1)
			} else { // regular fetch
				addToPushMap(r, 1)
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
