// +build go1.8

package link_test

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/romainmenke/pusher/link"
)

func TestHandler(t *testing.T) {

	var (
		server *httptest.Server
		rt     = &http.Transport{}
		client = &http.Client{Transport: rt}
	)

	h := link.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, ok := w.(http.Pusher)
		if !ok {
			t.Fatal("bad test, not a pusher")
		}
	}))

	server = httptest.NewUnstartedServer(h)
	server.TLS = &tls.Config{NextProtos: []string{"h2", "HTTP/1.1"}}
	server.StartTLS()

	{ // setup default config
		// fails because there is no server running at that address (but used to setup HTTP/2)
		client.Get("http://127.0.0.1:1/")
		if rt.TLSClientConfig == nil {
			rt.TLSClientConfig = &tls.Config{}
		}
		rt.TLSClientConfig.InsecureSkipVerify = true
	}
	client.Get(server.URL)
}

func BenchmarkHandlerA(b *testing.B) {

	var (
		server *httptest.Server
		rt     = &http.Transport{}
		client = &http.Client{Transport: rt}
	)

	h := link.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))

	server = httptest.NewUnstartedServer(h)
	server.TLS = &tls.Config{NextProtos: []string{"h2", "HTTP/1.1"}}
	server.StartTLS()

	{ // setup default config
		// fails because there is no server running at that address (but used to setup HTTP/2)
		client.Get("http://127.0.0.1:1/")
		if rt.TLSClientConfig == nil {
			rt.TLSClientConfig = &tls.Config{}
		}
		rt.TLSClientConfig.InsecureSkipVerify = true
	}

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			client.Get(server.URL)
		}
	})
}

func BenchmarkHandlerB(b *testing.B) {

	var (
		server *httptest.Server
		rt     = &http.Transport{}
		client = &http.Client{Transport: rt}
	)

	h := link.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			w.Header().Add("Link", "</css/stylesheet.css>; rel=preload; as=style;")
		default:
		}

		w.Write([]byte{})
	}))

	server = httptest.NewUnstartedServer(h)
	server.TLS = &tls.Config{NextProtos: []string{"h2", "HTTP/1.1"}}
	server.StartTLS()

	{ // setup default config
		// fails because there is no server running at that address (but used to setup HTTP/2)
		client.Get("http://127.0.0.1:1/")
		if rt.TLSClientConfig == nil {
			rt.TLSClientConfig = &tls.Config{}
		}
		rt.TLSClientConfig.InsecureSkipVerify = true
	}

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			client.Get(server.URL)
		}
	})
}

func BenchmarkDefaultHandlerA(b *testing.B) {

	var (
		server *httptest.Server
		rt     = &http.Transport{}
		client = &http.Client{Transport: rt}
	)

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})

	server = httptest.NewUnstartedServer(h)
	server.TLS = &tls.Config{NextProtos: []string{"h2", "HTTP/1.1"}}
	server.StartTLS()

	{ // setup default config
		// fails because there is no server running at that address (but used to setup HTTP/2)
		client.Get("http://127.0.0.1:1/")
		if rt.TLSClientConfig == nil {
			rt.TLSClientConfig = &tls.Config{}
		}
		rt.TLSClientConfig.InsecureSkipVerify = true
	}

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			client.Get(server.URL)
		}
	})
}

func BenchmarkDefaultHandlerB(b *testing.B) {

	var (
		server *httptest.Server
		rt     = &http.Transport{}
		client = &http.Client{Transport: rt}
	)

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			w.Header().Add("Link", "</css/stylesheet.css>; rel=preload; as=style;")
		default:
		}

		w.Write([]byte{})
	})

	server = httptest.NewUnstartedServer(h)
	server.TLS = &tls.Config{NextProtos: []string{"h2", "HTTP/1.1"}}
	server.StartTLS()

	{ // setup default config
		// fails because there is no server running at that address (but used to setup HTTP/2)
		client.Get("http://127.0.0.1:1/")
		if rt.TLSClientConfig == nil {
			rt.TLSClientConfig = &tls.Config{}
		}
		rt.TLSClientConfig.InsecureSkipVerify = true
	}

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			client.Get(server.URL)
		}
	})
}

func BenchmarkCanPush(b *testing.B) {

	request := &http.Request{
		Method:     "GET",
		ProtoMajor: 2,
		Header:     http.Header{},
	}

	var writer http.ResponseWriter

	writer = &testWriter{
		&httptest.ResponseRecorder{},
	}

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if link.CanPush(writer, request) {
				continue
			}
		}
	})
}

type testWriter struct {
	http.ResponseWriter
}

func (w *testWriter) Push(target string, opts *http.PushOptions) error {
	return nil
}
