package link

import (
	"log"
	"net/http"
)

// HandleFunc wraps an http.HandlerFunc with H2 Push functionality.
func HandleFunc(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if !CanPush(r) {
			handler(w, r)
			return
		}

		var rw = responseWriter(w)
		handler(rw, r)

	}
}

// CanPush checks if the Request is Pushable.
func CanPush(r *http.Request) bool {
	if r.Method != "GET" {
		return false
	}

	if r.Header.Get("Go-H2-Push") != "" {
		return false
	}

	return true
}

// InitiatePush parses Link Headers of a response to generate Push Frames.
func InitiatePush(header http.Header, pusher http.Pusher) { // 0 allocs

	linkHeaders, ok := header["Link"]
	if !ok {
		return
	}

	toPush, toLink := splitLinkHeadersAndParse(linkHeaders)

	for _, link := range toPush {
		if link == "" {
			continue
		}

		err := pusher.Push(link, &http.PushOptions{
			Header: http.Header{
				"Go-H2-Push": []string{link},
			},
		})
		if err != nil {
			log.Println(err)
		}
	}

	header["Link"] = toLink
	header["Go-H2-Pushed"] = toPush

}
