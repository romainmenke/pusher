package link

import (
	"net/http"
	"net/url"
	"strings"
)

type pushHandler struct {
	handlerFunc http.HandlerFunc
}

func (h *pushHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handlerFunc(w, r)
}

// Handler wraps a http.Handler.
func Handler(handler http.Handler) http.Handler {
	return &pushHandler{
		handlerFunc: newPushHandlerFunc(handler.ServeHTTP),
	}
}

// HandlerFunc wraps a http.HandlerFunc.
func HandlerFunc(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return newPushHandlerFunc(handlerFunc)
}

func newPushHandlerFunc(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler(newPusher(w), r)
	})
}

type pusher struct {
	writer http.ResponseWriter
	header http.Header
	status int
}

func newPusher(writer http.ResponseWriter) *pusher {
	return &pusher{writer: writer, header: make(http.Header)}
}

func (p *pusher) WriteHeader(rc int) {
	p.status = rc
}

func (p *pusher) Write(b []byte) (int, error) {
	p.Push()
	p.writer.WriteHeader(p.status)
	return p.writer.Write(b)
}

func (p *pusher) Push() {

	var (
		pusher     http.Pusher
		ok         bool
		linkHeader []string
	)

	for k, v := range p.Header() {
		if strings.ToLower(k) != "link" {
			p.writer.Header()[k] = v
			continue
		}
		linkHeader = v
	}

	pusher, ok = p.writer.(http.Pusher)
	if !ok || len(linkHeader) == 0 {
		return
	}

	for _, link := range linkHeader {
		parsed := parseLinkHeader(link)
		if parsed == "" || isAbsolute(parsed) {
			p.writer.Header().Add("link", link)
			continue
		}

		p.writer.Header().Add("Go-H2-Pushed", link)
		pusher.Push(parsed, nil)
	}

	return

}

func isAbsolute(p string) bool {
	u, err := url.Parse(p)
	if err != nil {
		return false
	}

	return u.IsAbs()
}

func (p *pusher) Header() http.Header {
	return p.header
}
