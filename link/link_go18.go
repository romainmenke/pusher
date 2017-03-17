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

		var rw = &responseWriter{ResponseWriter: w, request: r}
		handler.ServeHTTP(rw, r)

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

	if r.Header.Get("Go-H2-Push") != "" {
		return false
	}

	return true
}

// InitiatePush parses Link Headers of a response to generate Push Frames.
func InitiatePush(referer string, requestHeader http.Header, header http.Header, pusher http.Pusher) { // 0 allocs

	linkHeaders, ok := header["Link"]
	if !ok {
		return
	}

	toPush, toLink := splitLinkHeadersAndParse(linkHeaders)

	for _, link := range toPush {
		if link == "" {
			continue
		}

		pHeader := http.Header{}

		for k, v := range requestHeader {
			pHeader[k] = v
		}
		pHeader.Set("Go-H2-Pusher", referer)
		pHeader.Set("Go-H2-Pushed", link)

		err := pusher.Push(link, &http.PushOptions{
			Header: pHeader,
		})
		if err != nil {
			log.Println(err)
		}
	}

	header["Link"] = toLink
	header["Go-H2-Pushed"] = toPush

}
