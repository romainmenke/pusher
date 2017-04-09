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

func handlerBenchmarkFactory(length int) func(b *testing.B) {
	return func(b *testing.B) {
		var (
			server *httptest.Server
			rt     = &http.Transport{}
			client = &http.Client{Transport: rt}
		)

		links := make([]string, length)
		for i := 0; i < length; i++ {
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
}

func defaultHandlerBenchmarkFactory(length int) func(b *testing.B) {
	return func(b *testing.B) {
		var (
			server *httptest.Server
			rt     = &http.Transport{}
			client = &http.Client{Transport: rt}
		)

		links := make([]string, length)
		for i := 0; i < length; i++ {
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
}

func BenchmarkHandler(b *testing.B) {
	handlerBenchmarkFactory(0)(b)
}

func BenchmarkDefaultHandler(b *testing.B) {
	defaultHandlerBenchmarkFactory(0)(b)
}

func BenchmarkHandler_10(b *testing.B) {
	handlerBenchmarkFactory(10)(b)
}

func BenchmarkDefaultHandler_10(b *testing.B) {
	defaultHandlerBenchmarkFactory(10)(b)
}

func BenchmarkHandler_100(b *testing.B) {
	handlerBenchmarkFactory(100)(b)
}

func BenchmarkDefaultHandler_100(b *testing.B) {
	defaultHandlerBenchmarkFactory(100)(b)
}

func BenchmarkHandler_1000(b *testing.B) {
	handlerBenchmarkFactory(1000)(b)
}
func BenchmarkDefaultHandler_1000(b *testing.B) {
	defaultHandlerBenchmarkFactory(1000)(b)
}

func BenchmarkHandler_WorstCase(b *testing.B) {
	var (
		server *httptest.Server
		rt     = &http.Transport{}
		client = &http.Client{Transport: rt}
	)

	header := ""
	charLimit := (2048 - len("</>; rel=preload;"))
	for y := 0; y < charLimit; y++ {
		header += "a"
	}
	header = fmt.Sprintf("</%s>; rel=preload;", header)

	links := make([]string, 64)
	for i := 0; i < 64; i++ {
		links[i] = header
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

func BenchmarkDefaultHandler_WorstCase(b *testing.B) {
	var (
		server *httptest.Server
		rt     = &http.Transport{}
		client = &http.Client{Transport: rt}
	)

	header := ""
	charLimit := (2048 - len("</>; rel=preload;"))
	for y := 0; y < charLimit; y++ {
		header += "a"
	}
	header = fmt.Sprintf("</%s>; rel=preload;", header)

	links := make([]string, 64)
	for i := 0; i < 64; i++ {
		links[i] = header
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

func TestCanPush(t *testing.T) {
	request := &http.Request{
		Method:     "GET",
		ProtoMajor: 2,
		Header:     http.Header{},
	}
	var writer http.ResponseWriter
	writer = &testWriter{
		&httptest.ResponseRecorder{},
	}

	if !link.CanPush(writer, request) {
		t.Fail()
	}
}

func TestCanPush_H1(t *testing.T) {
	request := &http.Request{
		Method:     "GET",
		ProtoMajor: 1,
		Header:     http.Header{},
	}
	var writer http.ResponseWriter
	writer = &testWriter{
		&httptest.ResponseRecorder{},
	}

	if link.CanPush(writer, request) {
		t.Fail()
	}
}

func TestCanPush_Forwarded(t *testing.T) {
	request := &http.Request{
		Method:     "GET",
		ProtoMajor: 2,
		Header: http.Header{
			"X-Forwarded-For": []string{"foo"},
		},
	}
	var writer http.ResponseWriter
	writer = &testWriter{
		&httptest.ResponseRecorder{},
	}

	if link.CanPush(writer, request) {
		t.Fail()
	}
}

func TestCanPush_NoPusher(t *testing.T) {
	request := &http.Request{
		Method:     "GET",
		ProtoMajor: 2,
		Header:     http.Header{},
	}
	var writer http.ResponseWriter
	writer = &httptest.ResponseRecorder{}

	if link.CanPush(writer, request) {
		t.Fail()
	}
}

func TestCanPush_NoGet(t *testing.T) {
	request := &http.Request{
		Method:     "POST",
		ProtoMajor: 2,
		Header:     http.Header{},
	}
	var writer http.ResponseWriter
	writer = &testWriter{
		&httptest.ResponseRecorder{},
	}

	if link.CanPush(writer, request) {
		t.Fail()
	}
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

	for n := 0; n < b.N; n++ {
		if link.CanPush(writer, request) {
			continue
		}
	}
}

type testWriter struct {
	http.ResponseWriter
}

func (w *testWriter) Push(target string, opts *http.PushOptions) error {
	return nil
}
