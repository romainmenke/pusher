package adaptive

import (
	"fmt"
	"net/http"
)

// Handler wraps an http.Handler.
// It will automatically generate Push Promises
func Handler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		setInitiatorForWriter(w, r)
		if isAssetGet(r) {
			if isPushedResource(r) { // pushed content
				addToPushMap(r, 0.1)
			} else { // regular fetch
				addToPushMap(r, 1)
			}
			handler.ServeHTTP(w, r)
			return
		}

		if p, ok := w.(http.Pusher); ok {
			resourceKey := resourceKeyFromRequest(r)
			readFromPushMap(resourceKey, func(path string) {
				opts := &http.PushOptions{}
				opts = setInitiatorForOptions(r, opts)
				err := p.Push(path, opts)
				if err != nil {
					fmt.Println(err)
				}
			})
		}

		handler.ServeHTTP(w, r)

	})
}
