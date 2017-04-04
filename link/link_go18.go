// +build go1.8

package link

import (
	"log"
	"net/http"
)

// Handler wraps an http.HandlerFunc with H2 Push functionality.
func Handler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if !CanPush(w, r) {
			handler.ServeHTTP(w, r)
			return
		}

		var rw = getResponseWriter(w, r)

		handler.ServeHTTP(rw, r)
		rw.close()
	})
}

// CanPush checks if the Request is Pushable.
func CanPush(w http.ResponseWriter, r *http.Request) bool {

	if r.Method != "GET" {
		return false
	}

	if r.ProtoMajor < 2 {
		return false
	}

	_, ok := w.(http.Pusher)
	if !ok {
		return false
	}

	if r.Header.Get("Go-H2-Pushed") != "" {
		return false
	}

	return true
}

// InitiatePush parses Link Headers of a response to generate Push Frames.
func InitiatePush(w *responseWriter) { // 0 allocs

	if w == nil {
		return
	}

	pusher, ok := w.ResponseWriter.(http.Pusher)
	if !ok {
		return
	}

	linkHeaders, ok := w.Header()["Link"]
	if !ok {
		return
	}

	toPush, toLink := splitLinkHeadersAndParse(linkHeaders)

	for _, link := range toPush {
		if link == "" {
			continue
		}

		pHeader := http.Header{}

		if w.request != nil {
			for k, v := range w.request.Header {
				pHeader[k] = v
			}
			pHeader.Set("Go-H2-Pusher", w.request.URL.Path)
		}

		pHeader.Set("Go-H2-Pushed", link)

		err := pusher.Push(link, &http.PushOptions{
			Header: pHeader,
		})
		if err != nil {
			log.Println(err)
		}
	}

	w.ResponseWriter.Header()["Link"] = toLink
	w.ResponseWriter.Header()["Go-H2-Pushed"] = toPush

}
