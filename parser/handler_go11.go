// +build !go1.8

package parser

import (
	"net/http"

	"github.com/romainmenke/pusher/common"
)

// Handler wraps an http.Handler reading the response body and setting Link Headers or generating Pushes
func Handler(handler http.Handler, options ...Option) http.Handler {

	s := &settings{}
	for _, o := range options {
		o(s)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		writer := w
		defer handler.ServeHTTP(w, r)

		if r.Method != http.MethodGet {
			return
		}

		if s.withCache {
			preloads := getFromCache(r.URL.RequestURI())
			if preloads != nil {

				for _, link := range preloads {
					w.Header().Add(common.Link, link.LinkHeader())
				}

				return
			}
		}

		// Get a responseWriter from the sync.Pool.
		var rw = getResponseWriter(s, w, r)
		// defer close the responseWriter.
		// This returns it to the sync.Pool and zeroes all values and pointers.
		defer rw.close()

		switch r.Proto {
		case protoHTTP11:
			writer = &responseWriterHTTP11{
				responseWriter: rw,
			}
		case protoHTTP11TLS:
			writer = &responseWriterHTTP11TLS{
				responseWriter: rw,
			}
		case protoHTTP20:
			writer = &responseWriterHTTP2{
				responseWriter: rw,
			}
		}

	})
}
