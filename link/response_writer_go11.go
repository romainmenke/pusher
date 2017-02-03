// +build !go1.8

package link

import (
	"log"
	"net/http"
)

type responseWriter struct {
	http.ResponseWriter
}

func (r *responseWriter) Write(b []byte) (int, error) {
	return r.ResponseWriter.Write(b)
}

func (r *responseWriter) WriteHeader(s int) {
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
