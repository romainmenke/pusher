/*
https://golang.org/pkg/net/http/#ResponseWriter
https://golang.org/pkg/net/http/#Hijacker
https://golang.org/pkg/net/http/#Flusher
https://golang.org/pkg/net/http/#CloseNotifier
https://golang.org/pkg/io/#ReaderFrom
*/

package link

import (
	"bufio"
	"errors"
	"io"
	"net"
	"net/http"
)

type FullResponseWriter interface {
	http.ResponseWriter
	http.Hijacker
	http.Flusher
	// http.Pusher
	http.CloseNotifier
	io.ReaderFrom
}

var _ FullResponseWriter = &pusher{}

func (p *pusher) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	hijacker, ok := p.writer.(http.Hijacker)
	if !ok {
		return nil, nil, errors.New("webserver doesn't support hijacking")
	}
	return hijacker.Hijack()
}

func (p *pusher) Flush() {
	flusher, ok := p.writer.(http.Flusher)
	if !ok {
		return
	}
	flusher.Flush()
}

func (p *pusher) CloseNotify() <-chan bool {
	closeNotifier, ok := p.writer.(http.CloseNotifier)
	if !ok {
		return make(<-chan bool)
	}
	return closeNotifier.CloseNotify()
}

func (p *pusher) ReadFrom(r io.Reader) (n int64, err error) {
	readerFrom, ok := p.writer.(io.ReaderFrom)
	if !ok {
		return 0, errors.New("webserver doesn't support ReadFrom")
	}
	return readerFrom.ReadFrom(r)
}
