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

		if s.withCache {
			preloads := getFromCache(r.URL.RequestURI())
			if preloads != nil {
				defer handler.ServeHTTP(w, r)

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

		var protoW http.ResponseWriter
		switch r.Proto {
		case protoHTTP11:
			protoW = &responseWriterHTTP11{
				responseWriter: rw,
			}
		case protoHTTP11TLS:
			protoW = &responseWriterHTTP11TLS{
				responseWriter: rw,
			}
		case protoHTTP20:
			protoW = &responseWriterHTTP2{
				responseWriter: rw,
			}
		}

		// handle.
		handler.ServeHTTP(protoW, r)

	})
}
