// +build !go1.8

package rules_test

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/romainmenke/pusher/rules"
)

func TestHandlerGet(t *testing.T) {

	var (
		server *httptest.Server
		rt     = &http.Transport{}
		client = &http.Client{Transport: rt}
	)

	reader := strings.NewReader(`/
</css/stylesheet.css>; rel=preload; as=style;
</fonts/CutiveMono-Regular.ttf>; rel=preload; as=font;
/broken_a>; rel=preload;
</broken_b>

`)

	h := rules.Handler(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		}),
		rules.ReaderOption(reader),
	)

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

func TestHandlerPost(t *testing.T) {

	var (
		server *httptest.Server
		rt     = &http.Transport{}
		client = &http.Client{Transport: rt}
	)

	reader := strings.NewReader(`/
</css/stylesheet.css>; rel=preload; as=style;
</fonts/CutiveMono-Regular.ttf>; rel=preload; as=font;
/broken_a>; rel=preload;
</broken_b>

`)

	h := rules.Handler(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		}),
		rules.ReaderOption(reader),
	)

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
	resp, err := client.Post(server.URL, "application/json", nil)
	if err != nil {
		t.Fail()
	}
	if resp == nil {
		t.Fail()
	}
}

func TestHandlerGetPushed(t *testing.T) {

	var (
		server *httptest.Server
		rt     = &http.Transport{}
		client = &http.Client{Transport: rt}
	)

	reader := strings.NewReader(`/
</css/stylesheet.css>; rel=preload; as=style;
</fonts/CutiveMono-Regular.ttf>; rel=preload; as=font;
/broken_a>; rel=preload;
</broken_b>

`)

	h := rules.Handler(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		}),
		rules.ReaderOption(reader),
	)

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
	resp, err := client.Get(server.URL + "/css/stylesheet.css")
	if err != nil {
		t.Fail()
	}
	if resp == nil {
		t.Fail()
	}
}
