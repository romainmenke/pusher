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
}

func newPusher(writer http.ResponseWriter) *pusher {
	return &pusher{writer: writer}
}

func (p *pusher) WriteHeader(rc int) {
	p.writer.WriteHeader(rc)
}

func (p *pusher) Write(b []byte) (int, error) {

	var (
		pusher     http.Pusher
		ok         bool
		linkHeader []string
	)

	pusher, ok = p.writer.(http.Pusher)
	if !ok {
		return p.writer.Write(b)
	}

	for k, v := range p.writer.Header() {
		if strings.ToLower(k) != "link" {
			continue
		}
		linkHeader = v
	}

	if len(linkHeader) == 0 {
		return p.writer.Write(b)
	}

	for _, link := range linkHeader {
		parsed := parseLinkHeader(link)
		if parsed == "" || isAbsolute(parsed) {
			continue
		}
		pusher.Push(parsed, nil)
	}

	return p.writer.Write(b)
}

func parseLinkHeader(h string) string {

	var path string

	components := strings.Split(h, ";")
	for _, component := range components {
		if strings.HasPrefix(component, "<") && strings.HasSuffix(component, ">") {
			path = component
			path = strings.TrimPrefix(path, "<")
			path = strings.TrimSuffix(path, ">")
			continue
		}

		subComponents := strings.Split(component, "=")
		if len(subComponents) > 0 && subComponents[0] == "nopush" {
			return ""
		}

		if len(subComponents) > 1 && subComponents[0] == "rel" && subComponents[1] == "preload" {
			return path
		}
	}

	return ""
}

func isAbsolute(p string) bool {
	u, err := url.Parse(p)
	if err != nil {
		return false
	}

	return u.IsAbs()
}

func (p *pusher) Header() http.Header {
	return p.writer.Header()
}
