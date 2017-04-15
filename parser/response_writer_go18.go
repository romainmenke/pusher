// +build go1.8

package parser

import (
	"bytes"
	"net/http"

	"github.com/romainmenke/pusher/common"

	"golang.org/x/net/html"
)

type responseWriter struct {
	http.ResponseWriter

	// Code is the HTTP response code set by WriteHeader.
	//
	// Note that if a Handler never calls WriteHeader or Write,
	// this might end up being 0, rather than the implicit
	// http.StatusOK. To get the implicit value, use the Result
	// method.
	statusCode int

	// Body is the buffer to which the Handler's Write calls are sent.
	// If nil, the Writes are silently discarded.
	body *bytes.Buffer

	// Flushed is whether the Handler called Flush.
	flushed bool

	request *http.Request
}

func newResponseWriter(w http.ResponseWriter, r *http.Request) *responseWriter {
	return &responseWriter{
		w,
		0,
		new(bytes.Buffer),
		false,
		r,
	}
}

func (w *responseWriter) ExtractLinks() []common.Preloadable {
	links := make(map[common.Preloadable]struct{})
	preloads := make(map[string]struct{})

	contentType := http.DetectContentType(w.body.Bytes())
	if contentType != "text/html; charset=utf-8" {
		return nil
	}

	path := w.request.URL.RequestURI()

	z := html.NewTokenizer(w.body)

TOKENIZER:
	for {
		tt := z.Next()

		var asset common.Preloadable

		switch tt {
		case html.ErrorToken:
			// End of the document, we're done
			break TOKENIZER
		case html.StartTagToken:

			t := z.Token()

			switch t.Data {
			case "link":

				var isPreload bool

				for _, attr := range t.Attr {
					switch attr.Key {
					case "rel":
						if attr.Val == "preload" {
							isPreload = true
						}
					case "nopush":
						continue TOKENIZER
					case "href":
						if common.IsAbsolute(attr.Val) || attr.Val == path {
							continue TOKENIZER
						}
						if isPreload {
							preloads[attr.Val] = struct{}{}
						} else {
							asset = common.CSS(attr.Val)
						}
					}
				}

			case "script":

				for _, attr := range t.Attr {
					switch attr.Key {
					case "rel":
						if attr.Val == "preload" {
							continue TOKENIZER
						}
					case "nopush":
						continue TOKENIZER
					case "src":
						if common.IsAbsolute(attr.Val) || attr.Val == path {
							continue TOKENIZER
						}
						asset = common.JS(attr.Val)
					}
				}

			case "img":

				for _, attr := range t.Attr {
					switch attr.Key {
					case "rel":
						if attr.Val == "preload" {
							continue TOKENIZER
						}
					case "nopush":
						continue TOKENIZER
					case "src":
						if common.IsAbsolute(attr.Val) || attr.Val == path {
							continue TOKENIZER
						}
						asset = common.Img(attr.Val)
					}
				}

			}

			if asset != nil {
				if _, found := preloads[asset.Path()]; !found {
					links[asset] = struct{}{}
					asset = nil
				}
			}

		}

	}

	var linkSlice []common.Preloadable
	if len(links) > 64 {
		linkSlice = make([]common.Preloadable, 64)
	} else {
		linkSlice = make([]common.Preloadable, len(links))
	}
	index := 0
	for key := range links {
		if index >= len(linkSlice) {
			break
		}
		linkSlice[index] = key
		index++
	}

	return linkSlice
}

// Header returns the response headers.
func (w *responseWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

// Write always succeeds and writes to rw.Body, if not nil.
func (w *responseWriter) Write(buf []byte) (int, error) {
	if w.statusCode == 0 {
		w.statusCode = 200
	}

	if w.body != nil {
		l := len(buf)
		if l > 1024 {
			l = 1024
		}
		w.body.Write(buf[:l])
	}

	return w.ResponseWriter.Write(buf)
}

// WriteHeader sets rw.Code. After it is called, changing rw.Header
// will not affect rw.HeaderMap.
func (w *responseWriter) WriteHeader(s int) {
	if w.statusCode == 0 {
		w.statusCode = s
	}

	w.ResponseWriter.WriteHeader(w.statusCode)
}

// Flush sets rw.Flushed to true.
func (w *responseWriter) Flush() {
	flusher, ok := w.ResponseWriter.(http.Flusher)
	if ok && flusher != nil {
		flusher.Flush()
	}
}
