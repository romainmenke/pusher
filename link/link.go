package link

import "net/http"

func CanPush(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		return false
	}

	if r.Header.Get("Go-H2-Push") != "" {
		return false
	}

	var ok bool
	_, ok = w.(http.Pusher)
	if !ok {
		return false
	}

	return true
}

// Push Sends Push Frames to the client for each link header found in the response headers.
func Push(header http.Header, pusher http.Pusher) { // 0 allocs

	linkHeaders, ok := header["Link"]
	if !ok {
		return
	}

	toPush, toLink := ByPushable(linkHeaders).Split()

	for _, link := range toPush {
		parsed := parseLinkHeader(link)
		if parsed == "" {
			continue
		}

		pusher.Push(parsed, nil)
	}

	header["Link"] = toLink
	header["Go-H2-Pushed"] = toPush

}
