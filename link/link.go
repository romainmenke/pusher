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

	toPush, toLink := splitLinkHeadersAndParse(linkHeaders)

	for _, link := range toPush {
		pusher.Push(link, nil)
	}

	header["Link"] = toLink
	header["Go-H2-Pushed"] = toPush

}

func HandlerFunc(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return newPushHandlerFunc(handlerFunc)
}

func newPushHandlerFunc(handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if !CanPush(w, r) {
			handler(w, r)
			return
		}

		var p pusher
		p = pusher{writer: w}
		handler(&p, r)

	})
}

type pusher struct {
	writer http.ResponseWriter
}

func (p *pusher) Header() http.Header {
	return p.writer.Header()
}

func (p *pusher) Write(b []byte) (int, error) {
	return p.writer.Write(b)
}

func (p *pusher) WriteHeader(rc int) {
	Push(p.Header(), p.writer.(http.Pusher))
	p.writer.WriteHeader(rc)
}
