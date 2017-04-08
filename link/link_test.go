// +build go1.8

package link_test

import (
	"crypto/tls"
	"fmt"
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

func BenchmarkHandler(b *testing.B) {

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

func BenchmarkHandler_10(b *testing.B) {

	var (
		server *httptest.Server
		rt     = &http.Transport{}
		client = &http.Client{Transport: rt}
	)

	links := make([]string, 10)
	for i := 0; i < 10; i++ {
		links[i] = fmt.Sprintf("</css/stylesheet-%d.css>; rel=preload; as=style;", i)
	}

	h := link.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			w.Header()["Link"] = links
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

func BenchmarkHandler_100(b *testing.B) {

	var (
		server *httptest.Server
		rt     = &http.Transport{}
		client = &http.Client{Transport: rt}
	)

	links := make([]string, 100)
	for i := 0; i < 100; i++ {
		links[i] = fmt.Sprintf("</css/stylesheet-%d.css>; rel=preload; as=style;", i)
	}

	h := link.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			w.Header()["Link"] = links
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

func BenchmarkHandler_1000(b *testing.B) {

	var (
		server *httptest.Server
		rt     = &http.Transport{}
		client = &http.Client{Transport: rt}
	)

	links := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		links[i] = fmt.Sprintf("</css/stylesheet-%d.css>; rel=preload; as=style;", i)
	}

	h := link.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			w.Header()["Link"] = links
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

func BenchmarkHandler_10000(b *testing.B) {

	var (
		server *httptest.Server
		rt     = &http.Transport{}
		client = &http.Client{Transport: rt}
	)

	links := make([]string, 10000)
	for i := 0; i < 10000; i++ {
		links[i] = fmt.Sprintf("</css/stylesheet-%d.css>; rel=preload; as=style;", i)
	}

	h := link.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			w.Header()["Link"] = links
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

func BenchmarkDefaultHandler(b *testing.B) {

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

func BenchmarkDefaultHandler_10(b *testing.B) {

	var (
		server *httptest.Server
		rt     = &http.Transport{}
		client = &http.Client{Transport: rt}
	)

	links := make([]string, 10)
	for i := 0; i < 10; i++ {
		links[i] = fmt.Sprintf("</css/stylesheet-%d.css>; rel=preload; as=style;", i)
	}

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			w.Header()["Link"] = links
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

func BenchmarkDefaultHandler_100(b *testing.B) {

	var (
		server *httptest.Server
		rt     = &http.Transport{}
		client = &http.Client{Transport: rt}
	)

	links := make([]string, 100)
	for i := 0; i < 100; i++ {
		links[i] = fmt.Sprintf("</css/stylesheet-%d.css>; rel=preload; as=style;", i)
	}

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			w.Header()["Link"] = links
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

func BenchmarkDefaultHandler_1000(b *testing.B) {

	var (
		server *httptest.Server
		rt     = &http.Transport{}
		client = &http.Client{Transport: rt}
	)

	links := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		links[i] = fmt.Sprintf("</css/stylesheet-%d.css>; rel=preload; as=style;", i)
	}

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			w.Header()["Link"] = links
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

func BenchmarkDefaultHandler_10000(b *testing.B) {

	var (
		server *httptest.Server
		rt     = &http.Transport{}
		client = &http.Client{Transport: rt}
	)

	links := make([]string, 10000)
	for i := 0; i < 10000; i++ {
		links[i] = fmt.Sprintf("</css/stylesheet-%d.css>; rel=preload; as=style;", i)
	}

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			w.Header()["Link"] = links
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
