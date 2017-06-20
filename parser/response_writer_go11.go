// +build !go1.8

package parser

import "github.com/romainmenke/pusher/common"

// Write always succeeds and writes to rw.Body, if not nil.
func (w *responseWriter) Write(buf []byte) (int, error) {
	if w.statusCode == 0 {
		w.statusCode = 200
	}

	if !w.headerWritten {
		w.headerWritten = true

		if w.body != nil {
			l := len(buf)
			if l > 1024 {
				l = 1024
			}
			w.body.Write(buf[:l])
		}

		links := w.extractLinks()
		for {
			link, more := <-links
			if more {
				w.Header().Add(common.Link, link.LinkHeader())
			} else {
				break
			}
		}
	}

	w.ResponseWriter.WriteHeader(w.statusCode)

	return w.ResponseWriter.Write(buf)
}

func (w *responseWriter) WriteString(s string) (n int, err error) {
	if w.statusCode == 0 {
		w.statusCode = 200
	}

	if !w.headerWritten {
		w.headerWritten = true

		if w.body != nil {
			l := len(s)
			if l > 1024 {
				l = 1024
			}
			w.body.WriteString(s[:l])
		}

		links := w.extractLinks()
		for {
			link, more := <-links
			if more {
				w.Header().Add(common.Link, link.LinkHeader())
			} else {
				break
			}
		}
	}

	w.ResponseWriter.WriteHeader(w.statusCode)

	stringWriter, ok := w.ResponseWriter.(common.StringWriter)
	if ok {
		return stringWriter.WriteString(s)
	}
	return w.ResponseWriter.Write([]byte(s))
}
