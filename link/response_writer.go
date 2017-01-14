package link

import (
	"bufio"
	"io"
	"log"
	"net"
	"net/http"
)

type responseWriter struct {
	writer     http.ResponseWriter
	statusCode int
}

func (r *responseWriter) Header() http.Header {
	return r.writer.Header()
}

func (r *responseWriter) Write(b []byte) (int, error) {
	if r.statusCode == 0 {
		r.statusCode = 200
	}

	return r.writer.Write(b)
}

func (r *responseWriter) WriteHeader(s int) {
	if r.statusCode == 0 {
		r.statusCode = s
	}

	if r.statusCode/100 == 2 {
		InitiatePush(r.Header(), r)
	}

	r.writer.WriteHeader(s)
}

func (r *responseWriter) Flush() {
	flusher, ok := r.writer.(http.Flusher)
	if ok && flusher != nil {
		flusher.Flush()
	} else {
		log.Printf("Failed flush(%T)", r.writer)
	}
}

func (r *responseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return r.writer.(http.Hijacker).Hijack()
}

func (r *responseWriter) CloseNotify() <-chan bool {
	return r.writer.(http.CloseNotifier).CloseNotify()
}

func (r *responseWriter) ReadFrom(reader io.Reader) (int64, error) {
	if r.statusCode == 0 {
		r.statusCode = 200
	}

	return r.writer.(io.ReaderFrom).ReadFrom(reader)
}

func (r *responseWriter) Push(target string, opts *http.PushOptions) error {
	pusher, ok := r.writer.(http.Pusher)
	if ok && pusher != nil {
		return pusher.Push(target, opts)
	}
	return http.ErrNotSupported
}
