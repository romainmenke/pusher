// +build go1.8

package link

import "net/http"

type responseWriter struct {
	http.ResponseWriter
	request    *http.Request
	statusCode int
}

func (w *responseWriter) reset() *responseWriter {
	w.request = nil
	w.ResponseWriter = nil
	w.statusCode = 0

	return w
}

func (w *responseWriter) close() {
	w.reset()
	writerPool.Put(w)
}

func (w *responseWriter) Write(b []byte) (int, error) {
	if w.statusCode == 0 {
		w.statusCode = 200
	}

	return w.ResponseWriter.Write(b)
}

func (w *responseWriter) WriteHeader(s int) {
	if w.statusCode == 0 {
		w.statusCode = s
	}

	if w.statusCode/100 == 2 {
		InitiatePush(w)
	}

	w.ResponseWriter.WriteHeader(s)
}

func (w *responseWriter) Flush() {
	flusher, ok := w.ResponseWriter.(http.Flusher)
	if ok && flusher != nil {
		flusher.Flush()
	}
}

func (w *responseWriter) CloseNotify() <-chan bool {
	return w.ResponseWriter.(http.CloseNotifier).CloseNotify()
}

func (w *responseWriter) Push(target string, opts *http.PushOptions) error {
	pusher, ok := w.ResponseWriter.(http.Pusher)
	if ok && pusher != nil {
		return pusher.Push(target, opts)
	}
	return http.ErrNotSupported
}
