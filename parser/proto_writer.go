package parser

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
)

type responseWriterHTTP11 struct {
	*responseWriter
}

func (w *responseWriterHTTP11) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	if hj, ok := w.ResponseWriter.(http.Hijacker); ok {
		return hj.Hijack()
	}
	return nil, nil, fmt.Errorf("http.Hijacker interface is not supported")
}

func (w *responseWriterHTTP11) ReadFrom(reader io.Reader) (int64, error) {

	if !w.headerWritten {

		if w.statusCode == 0 {
			w.statusCode = 200
		}

		limitReader := io.LimitReader(reader, 1024)
		b, err := ioutil.ReadAll(limitReader)
		if err != nil {
			return 0, err
		}

		w.Write(b)

	}

	return w.ResponseWriter.(io.ReaderFrom).ReadFrom(reader)
}

type responseWriterHTTP11TLS struct {
	*responseWriter
}

func (w *responseWriterHTTP11TLS) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	if hj, ok := w.ResponseWriter.(http.Hijacker); ok {
		return hj.Hijack()
	}
	return nil, nil, fmt.Errorf("http.Hijacker interface is not supported")
}

func (w *responseWriterHTTP11TLS) ReadFrom(reader io.Reader) (int64, error) {
	if !w.headerWritten {

		if w.statusCode == 0 {
			w.statusCode = 200
		}

		limitReader := io.LimitReader(reader, 1024)
		b, err := ioutil.ReadAll(limitReader)
		if err != nil {
			return 0, err
		}

		w.Write(b)

	}

	return w.ResponseWriter.(io.ReaderFrom).ReadFrom(reader)
}

type responseWriterHTTP2 struct {
	*responseWriter
}
