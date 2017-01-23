package link

import (
	"log"
	"net/http"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (r *responseWriter) Write(b []byte) (int, error) {
	if r.statusCode == 0 {
		r.statusCode = 200
	}

	return r.ResponseWriter.Write(b)
}

func (r *responseWriter) WriteHeader(s int) {
	if r.statusCode == 0 {
		r.statusCode = s
	}

	if r.statusCode/100 == 2 {
		InitiatePush(r.Header(), r)
	}

	r.ResponseWriter.WriteHeader(s)
}

func (r *responseWriter) Flush() {
	flusher, ok := r.ResponseWriter.(http.Flusher)
	if ok && flusher != nil {
		flusher.Flush()
	} else {
		log.Printf("Failed flush(%T)", r.ResponseWriter)
	}
}

func (r *responseWriter) CloseNotify() <-chan bool {
	return r.ResponseWriter.(http.CloseNotifier).CloseNotify()
}

func (r *responseWriter) Push(target string, opts *http.PushOptions) error {
	pusher, ok := r.ResponseWriter.(http.Pusher)
	if ok && pusher != nil {
		return pusher.Push(target, opts)
	}
	return http.ErrNotSupported
}
