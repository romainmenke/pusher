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

func TestHandlerGet(t *testing.T) {

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
	resp, err := client.Get(server.URL)
	if err != nil {
		t.Fail()
	}
	if resp == nil {
		t.Fail()
	}
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
				resp, err := client.Get(server.URL)
				if err != nil {
					b.Fail()
				}
				if resp == nil {
					b.Fail()
				}
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
				resp, err := client.Get(server.URL)
				if err != nil {
					b.Fail()
				}
				if resp == nil {
					b.Fail()
				}
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
	charLimit := (1024 - len("</>; rel=preload;"))
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
			resp, err := client.Get(server.URL)
			if err != nil {
				b.Fail()
			}
			if resp == nil {
				b.Fail()
			}
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
	charLimit := (1024 - len("</>; rel=preload;"))
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
			resp, err := client.Get(server.URL)
			if err != nil {
				b.Fail()
			}
			if resp == nil {
				b.Fail()
			}
		}
	})
}
