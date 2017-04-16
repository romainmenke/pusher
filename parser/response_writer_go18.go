// +build go1.8

package parser

import (
	"net/http"

	"github.com/romainmenke/pusher/common"
)

// Write always succeeds and writes to rw.Body, if not nil.
func (w *responseWriter) Write(buf []byte) (int, error) {
	if w.statusCode == 0 {
		w.statusCode = 200
	}

	if !w.headerWritter {
		w.headerWritter = true

		if w.body != nil {
			l := len(buf)
			if l > 1024 {
				l = 1024
			}
			w.body.Write(buf[:l])
		}

		p := w.extractLinks()

		if pusher, ok := w.ResponseWriter.(http.Pusher); ok && w.request.Header.Get(common.XForwardedFor) == "" {
			for _, l := range p {
				pusher.Push(l.Path(), &http.PushOptions{
					Header: w.request.Header,
				})
			}
		} else {
			for _, l := range p {
				w.Header().Add(common.Link, l.LinkHeader())
			}
		}

		w.ResponseWriter.WriteHeader(w.statusCode)
	}

	return w.ResponseWriter.Write(buf)
}
