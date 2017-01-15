package link

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
)

func responseWriter(rw http.ResponseWriter) http.ResponseWriter {
	var w http.ResponseWriter

	switch rw.(type) {
	case io.ReaderFrom:
		w = &reHiFlPuClReResponseWriter{ResponseWriter: w}
		return w
	default:
		w = &reHiFlPuClResponseWriter{ResponseWriter: rw}
		return w
	}
}

type reHiFlPuClRe interface {
	http.ResponseWriter
	http.Hijacker
	http.Flusher
	http.Pusher
	http.CloseNotifier
	io.ReaderFrom
}

type reHiFlPuCl interface {
	http.ResponseWriter
	http.Hijacker
	http.Flusher
	http.Pusher
	http.CloseNotifier
}

type reHiFlPuClReResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

var _ reHiFlPuClRe = &reHiFlPuClReResponseWriter{}

func (r *reHiFlPuClReResponseWriter) Write(b []byte) (int, error) {
	if r.statusCode == 0 {
		r.statusCode = 200
	}

	return r.ResponseWriter.Write(b)
}

func (r *reHiFlPuClReResponseWriter) WriteHeader(s int) {
	if r.statusCode == 0 {
		r.statusCode = s
	}

	if r.statusCode/100 == 2 {
		InitiatePush(r.Header(), r)
	}

	r.ResponseWriter.WriteHeader(s)
}

func (r *reHiFlPuClReResponseWriter) Flush() {
	flusher, ok := r.ResponseWriter.(http.Flusher)
	if ok && flusher != nil {
		flusher.Flush()
	} else {
		log.Printf("Failed flush(%T)", r.ResponseWriter)
	}
}

func (r *reHiFlPuClReResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	if hj, ok := r.ResponseWriter.(http.Hijacker); ok {
		return hj.Hijack()
	}
	return nil, nil, fmt.Errorf("http.Hijacker interface is not supported")
}

func (r *reHiFlPuClReResponseWriter) CloseNotify() <-chan bool {
	return r.ResponseWriter.(http.CloseNotifier).CloseNotify()
}

func (r *reHiFlPuClReResponseWriter) ReadFrom(reader io.Reader) (int64, error) {
	if r.statusCode == 0 {
		r.statusCode = 200
	}

	return r.ResponseWriter.(io.ReaderFrom).ReadFrom(reader)
}

func (r *reHiFlPuClReResponseWriter) Push(target string, opts *http.PushOptions) error {
	pusher, ok := r.ResponseWriter.(http.Pusher)
	if ok && pusher != nil {
		return pusher.Push(target, opts)
	}
	return http.ErrNotSupported
}

type reHiFlPuClResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

var _ reHiFlPuCl = &reHiFlPuClResponseWriter{}

func (r *reHiFlPuClResponseWriter) Write(b []byte) (int, error) {
	if r.statusCode == 0 {
		r.statusCode = 200
	}

	return r.ResponseWriter.Write(b)
}

func (r *reHiFlPuClResponseWriter) WriteHeader(s int) {
	if r.statusCode == 0 {
		r.statusCode = s
	}

	if r.statusCode/100 == 2 {
		InitiatePush(r.Header(), r)
	}

	r.ResponseWriter.WriteHeader(s)
}

func (r *reHiFlPuClResponseWriter) Flush() {
	flusher, ok := r.ResponseWriter.(http.Flusher)
	if ok && flusher != nil {
		flusher.Flush()
	} else {
		log.Printf("Failed flush(%T)", r.ResponseWriter)
	}
}

func (r *reHiFlPuClResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	if hj, ok := r.ResponseWriter.(http.Hijacker); ok {
		return hj.Hijack()
	}
	return nil, nil, fmt.Errorf("http.Hijacker interface is not supported")
}

func (r *reHiFlPuClResponseWriter) CloseNotify() <-chan bool {
	return r.ResponseWriter.(http.CloseNotifier).CloseNotify()
}

func (r *reHiFlPuClResponseWriter) Push(target string, opts *http.PushOptions) error {
	pusher, ok := r.ResponseWriter.(http.Pusher)
	if ok && pusher != nil {
		return pusher.Push(target, opts)
	}
	return http.ErrNotSupported
}
