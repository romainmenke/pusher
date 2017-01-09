package link

import "net/http"

type pushHandler struct {
	handlerFunc http.HandlerFunc
}

func (h *pushHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handlerFunc(w, r)
}

// Handler wraps a http.Handler.
func Handler(handler http.Handler) http.Handler {
	h := pushHandler{
		handlerFunc: newPushHandlerFunc(handler.ServeHTTP),
	}
	return &h
}

// HandlerFunc wraps a http.HandlerFunc.
func HandlerFunc(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return newPushHandlerFunc(handlerFunc)
}

func newPushHandlerFunc(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// only push on "GET" requests
		if r.Method != "GET" {
			handler(w, r)
			return
		}

		// if the client does not support H2 Push, abort as early as possible
		var ok bool
		_, ok = w.(http.Pusher)
		if !ok {
			handler(w, r)
			return
		}

		var p pusher
		var header http.Header
		header = make(http.Header)
		p = pusher{writer: w, header: header}
		handler(&p, r)
	})
}

// A ResponseWriter struct is used by an HTTP handler to
// construct an HTTP response with Push Frames for link headers.
//
// A ResponseWriter may not be used after the Handler.ServeHTTP method
// has returned.
type pusher struct {
	writer http.ResponseWriter
	header http.Header
	status int
}

// Header returns the header map that will be sent by
// WriteHeader. The Header map also is the mechanism with which
// Handlers can set HTTP trailers.
//
// Changing the header map after a call to WriteHeader (or
// Write) has no effect unless the modified headers are
// trailers.
//
// There are two ways to set Trailers. The preferred way is to
// predeclare in the headers which trailers you will later
// send by setting the "Trailer" header to the names of the
// trailer keys which will come later. In this case, those
// keys of the Header map are treated as if they were
// trailers. See the example. The second way, for trailer
// keys not known to the Handler until after the first Write,
// is to prefix the Header map keys with the TrailerPrefix
// constant value. See TrailerPrefix.
//
// To suppress implicit response headers (such as "Date"), set
// their value to nil.
func (p *pusher) Header() http.Header {
	return p.header
}

// Write writes the data to the connection as part of an HTTP reply.
//
// If WriteHeader has not yet been called, Write calls
// WriteHeader(http.StatusOK) before writing the data. If the Header
// does not contain a Content-Type line, Write adds a Content-Type set
// to the result of passing the initial 512 bytes of written data to
// DetectContentType.
//
// Depending on the HTTP protocol version and the client, calling
// Write or WriteHeader may prevent future reads on the
// Request.Body. For HTTP/1.x requests, handlers should read any
// needed request body data before writing the response. Once the
// headers have been flushed (due to either an explicit Flusher.Flush
// call or writing enough data to trigger a flush), the request body
// may be unavailable. For HTTP/2 requests, the Go HTTP server permits
// handlers to continue to read the request body while concurrently
// writing the response. However, such behavior may not be supported
// by all HTTP/2 clients. Handlers should read before writing if
// possible to maximize compatibility.
func (p *pusher) Write(b []byte) (int, error) {
	return p.writer.Write(b)
}

// WriteHeader sends an HTTP response header with status code.
// If WriteHeader is not called explicitly, the first call to Write
// will trigger an implicit WriteHeader(http.StatusOK).
// Thus explicit calls to WriteHeader are mainly used to
// send error codes.
func (p *pusher) WriteHeader(rc int) {
	p.Push()
	p.writer.WriteHeader(rc)
}

// Push Sends Push Frames to the client for each link header found in the response headers.
func (p *pusher) Push() {

	var (
		pusher http.Pusher
	)

	pusher = p.writer.(http.Pusher)

	for k, v := range p.Header() {
		if k != "Link" {
			p.writer.Header()[k] = v
			continue
		}

		for _, link := range v {
			parsed := parseLinkHeader(link)
			if parsed == "" {
				p.writer.Header().Add("Link", link)
				continue
			}

			p.writer.Header().Add("Go-H2-Pushed", link)
			pusher.Push(parsed, nil)
		}
	}

	return

}

type LinkHeaderSlice []string

func (s LinkHeaderSlice) Len() int {
	return len(s)
}
func (s LinkHeaderSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s LinkHeaderSlice) Less(i, j int) bool {
	parsedI := parseLinkHeader(s[i])
	parsedJ := parseLinkHeader(s[i])

	if parsedI == "" && parsedJ == "" {
		return false
	}

	if parsedI != "" && parsedJ != "" {
		return len(s[i]) < len(s[j])
	}

	if parsedI != "" {
		return true
	}

	return false
}
