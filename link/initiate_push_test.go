// +build go1.8

package link

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/romainmenke/pusher/common"
)

func TestInitatePush(t *testing.T) {
	request := &http.Request{
		Method:     "GET",
		ProtoMajor: 2,
		Header: http.Header{
			"Accept-Encoding": []string{"gzip"},
		},
	}
	testW := &testWriter{
		[]string{},
		nil,
		httptest.NewRecorder(),
	}

	writer := &responseWriter{
		testW,
		request,
		0,
		false,
	}

	writer.Header()["Link"] = []string{"</style.css>; rel=preload;"}

	InitiatePush(writer)

	if len(testW.pushed) != 1 || testW.pushed[0] != "/style.css" {
		t.Fatal("bad push", testW.pushed)
	}

	if testW.options == nil || testW.options.Header == nil {
		t.Fatal("bad options")
	}

	if len(testW.options.Header["Accept-Encoding"]) != 1 || testW.options.Header["Accept-Encoding"][0] != "gzip" {
		t.Fatal("bad options header")
	}
}

func TestInitatePush_AbsoluteLink(t *testing.T) {
	request := &http.Request{
		Method:     "GET",
		ProtoMajor: 2,
		Header: http.Header{
			"Foo": []string{"Baz"},
		},
	}
	testW := &testWriter{
		[]string{},
		nil,
		httptest.NewRecorder(),
	}

	writer := &responseWriter{
		testW,
		request,
		0,
		false,
	}

	writer.Header()["Link"] = []string{"<www.site.com/style.css>; rel=preload;"}

	InitiatePush(writer)

	if len(testW.pushed) != 0 {
		t.Fatal("bad push")
	}

	if testW.options != nil {
		t.Fatal("bad options")
	}
}

func TestInitatePush_Mixed(t *testing.T) {
	request := &http.Request{
		Method:     "GET",
		ProtoMajor: 2,
		Header: http.Header{
			"Accept-Encoding": []string{"gzip"},
		},
	}
	testW := &testWriter{
		[]string{},
		nil,
		httptest.NewRecorder(),
	}

	writer := &responseWriter{
		testW,
		request,
		0,
		false,
	}

	writer.Header()["Link"] = []string{
		"</font>;",
		"</style.css>; rel=preload;",
		"</bundle.js>; rel=preload; nopush;",
		"<www.site.com/image.jpg>; rel=preload;",
		"asdfghjkjncbdfgfhdgfhgfh; sfdg; asjdfbklfnjkjksdf",
	}

	InitiatePush(writer)

	if len(testW.pushed) != 1 || testW.pushed[0] != "/style.css" {
		t.Fatal("bad push")
	}

	if testW.options == nil || testW.options.Header == nil {
		t.Fatal("bad options")
	}

	if len(testW.options.Header["Accept-Encoding"]) != 1 || testW.options.Header["Accept-Encoding"][0] != "gzip" {
		t.Fatal("bad options header")
	}

	if writer.Header()["Go-H2-Pushed"][0] != "</style.css>; rel=preload;" {
		t.Fatal("bad push header")
	}

	found := 0
	for _, v := range writer.Header()["Link"] {
		switch v {
		case "</font>;",
			"</bundle.js>; rel=preload; nopush;",
			"<www.site.com/image.jpg>; rel=preload;",
			"asdfghjkjncbdfgfhdgfhgfh; sfdg; asjdfbklfnjkjksdf":
			found++
		}
	}
	if found != 4 {
		t.Fatal("bad link header")
	}
}

func TestInitiatePushLinkLimit(t *testing.T) {
	request := &http.Request{
		Method:     "GET",
		ProtoMajor: 2,
		Header:     http.Header{},
	}

	writer := getResponseWriter(
		&testWriter{
			[]string{},
			nil,
			httptest.NewRecorder(),
		},
		request,
	)
	defer writer.close()

	writer.Header()[common.Link] = []string{}

	for i := 0; i < 80; i++ {
		writer.Header()[common.Link] = append(writer.Header()[common.Link], fmt.Sprintf("</css/stylesheet-%d.css>; rel=preload; as=style;", i))
	}

	InitiatePush(writer)

	link, ok := writer.Header()["Link"]
	if !ok {
		t.Fatal("missing link header")
	}
	if len(link) != 16 {
		t.Fatal("bad link header", len(link))
	}

	pushed, ok := writer.Header()["Go-H2-Pushed"]
	if !ok {
		t.Fatal("missing push header")
	}
	if len(pushed) != 64 {
		t.Fatal("bad push header", len(pushed))
	}

}

type testWriter struct {
	pushed  []string
	options *http.PushOptions
	http.ResponseWriter
}

func (w *testWriter) Push(target string, opts *http.PushOptions) error {
	w.pushed = append(w.pushed, target)
	w.options = opts
	return nil
}

func TestInitiatePushNoPusher(t *testing.T) {
	request := &http.Request{
		Method:     "GET",
		ProtoMajor: 2,
		Header:     http.Header{},
	}

	writer := getResponseWriter(
		httptest.NewRecorder(),
		request,
	)

	writer.Header()[common.Link] = []string{
		"</css/stylesheet-1.css>; rel=preload; as=style;",
		"</css/stylesheet-2.css>; rel=preload; as=style;",
		"</css/stylesheet-3.css>; rel=preload; as=style;",
	}

	InitiatePush(writer)

	link, ok := writer.Header()["Link"]
	if !ok {
		t.Fatal("missing link header")
	}
	if len(link) != 3 {
		t.Fatal("bad link header")
	}

	writer.close()

	// nil request

	writer = getResponseWriter(
		httptest.NewRecorder(),
		nil,
	)

	writer.Header()[common.Link] = []string{
		"</css/stylesheet-1.css>; rel=preload; as=style;",
		"</css/stylesheet-2.css>; rel=preload; as=style;",
		"</css/stylesheet-3.css>; rel=preload; as=style;",
	}

	InitiatePush(writer)

	link, ok = writer.Header()["Link"]
	if !ok {
		t.Fatal("missing link header")
	}
	if len(link) != 3 {
		t.Fatal("bad link header")
	}

	writer.close()

	// nil writer

	InitiatePush(nil)
}

func TestInitiatePushRandomErr(t *testing.T) {
	request := &http.Request{
		Method:     "GET",
		ProtoMajor: 2,
		Header:     http.Header{},
	}

	writer := getResponseWriter(
		&testWriterErr{
			0,
			httptest.NewRecorder(),
		},
		request,
	)
	defer writer.close()

	writer.Header()[common.Link] = []string{
		"</css/stylesheet-1.css>; rel=preload; as=style;",
		"</css/stylesheet-2.css>; rel=preload; as=style;",
		"</css/stylesheet-3.css>; rel=preload; as=style;",
	}

	InitiatePush(writer)

	pushed, ok := writer.Header()["Go-H2-Pushed"]
	if !ok {
		t.Fatal("missing push header")
	}
	if len(pushed) != 2 && pushed[0] != "</css/stylesheet-1.css>; rel=preload; as=style;" && pushed[1] != "</css/stylesheet-3.css>; rel=preload; as=style;" {
		t.Fatal("bad push header")
	}

	link, ok := writer.Header()["Link"]
	if !ok {
		t.Fatal("missing link header")
	}
	if len(link) != 1 && link[0] != "</css/stylesheet-2.css>; rel=preload; as=style;" {
		t.Fatal("bad link header")
	}
}

func TestInitiatePushRecursiveErr(t *testing.T) {
	request := &http.Request{
		Method:     "GET",
		ProtoMajor: 2,
		Header:     http.Header{},
	}

	writer := getResponseWriter(
		&testWriterRecursiveErr{
			0,
			httptest.NewRecorder(),
		},
		request,
	)
	defer writer.close()

	writer.Header()[common.Link] = []string{
		"</css/stylesheet-1.css>; rel=preload; as=style;",
		"</css/stylesheet-2.css>; rel=preload; as=style;",
		"</css/stylesheet-3.css>; rel=preload; as=style;",
	}

	InitiatePush(writer)

	pushed, ok := writer.Header()["Go-H2-Pushed"]
	if !ok {
		t.Fatal("missing push header")
	}
	if len(pushed) != 1 && pushed[0] != "</css/stylesheet-1.css>; rel=preload; as=style;" {
		t.Fatal("bad push header")
	}

	link, ok := writer.Header()["Link"]
	if !ok {
		t.Fatal("missing link header")
	}
	if len(link) != 2 && link[0] != "</css/stylesheet-2.css>; rel=preload; as=style;" && link[1] != "</css/stylesheet-3.css>; rel=preload; as=style;" {
		t.Fatal("bad link header")
	}
}

func TestInitiatePushMaxStreamsErr(t *testing.T) {
	request := &http.Request{
		Method:     "GET",
		ProtoMajor: 2,
		Header:     http.Header{},
	}

	writer := getResponseWriter(
		&testWriterConcurrentStreamsErr{
			0,
			httptest.NewRecorder(),
		},
		request,
	)
	defer writer.close()

	writer.Header()[common.Link] = []string{
		"</css/stylesheet-1.css>; rel=preload; as=style;",
		"</css/stylesheet-2.css>; rel=preload; as=style;",
		"</css/stylesheet-3.css>; rel=preload; as=style;",
	}

	InitiatePush(writer)

	pushed, ok := writer.Header()["Go-H2-Pushed"]
	if !ok {
		t.Fatal("missing push header")
	}
	if len(pushed) != 1 && pushed[0] != "</css/stylesheet-1.css>; rel=preload; as=style;" {
		t.Fatal("bad push header")
	}

	link, ok := writer.Header()["Link"]
	if !ok {
		t.Fatal("missing link header")
	}
	if len(link) != 2 && link[0] != "</css/stylesheet-2.css>; rel=preload; as=style;" && link[1] != "</css/stylesheet-3.css>; rel=preload; as=style;" {
		t.Fatal("bad link header")
	}
}

type testWriterErr struct {
	pushes int
	http.ResponseWriter
}

func (w *testWriterErr) Push(target string, opts *http.PushOptions) error {
	if w.pushes == 1 {
		w.pushes++
		return errors.New("random err")
	}
	w.pushes++
	return nil
}

type testWriterRecursiveErr struct {
	pushes int
	http.ResponseWriter
}

func (w *testWriterRecursiveErr) Push(target string, opts *http.PushOptions) error {
	if w.pushes == 1 {
		w.pushes++
		return errors.New("http2: recursive push not allowed")
	}
	w.pushes++
	return nil
}

type testWriterConcurrentStreamsErr struct {
	pushes int
	http.ResponseWriter
}

func (w *testWriterConcurrentStreamsErr) Push(target string, opts *http.PushOptions) error {
	if w.pushes == 1 {
		w.pushes++
		return errors.New("http2: push would exceed peer's SETTINGS_MAX_CONCURRENT_STREAMS")
	}
	w.pushes++
	return nil
}
