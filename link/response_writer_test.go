// +build go1.8

package link

import (
	"net/http"
	"net/http/httptest"
	"testing"

	httpmiddlewarevet "github.com/fd/httpmiddlewarevet"
)

func TestMiddleware(t *testing.T) {
	httpmiddlewarevet.Vet(t, func(h http.Handler) http.Handler {
		return Handler(h)
	})
}

func TestInitatePush(t *testing.T) {
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
		&httptest.ResponseRecorder{},
	}

	writer := &responseWriter{
		testW,
		request,
		0,
	}

	writer.Header()["Link"] = []string{"</style.css>; rel=preload;"}

	InitiatePush(writer)

	if len(testW.pushed) != 1 || testW.pushed[0] != "/style.css" {
		t.Fatal("bad push")
	}

	if testW.options == nil || testW.options.Header == nil {
		t.Fatal("bad options")
	}

	if len(testW.options.Header["Foo"]) != 1 || testW.options.Header["Foo"][0] != "Baz" {
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
		&httptest.ResponseRecorder{},
	}

	writer := &responseWriter{
		testW,
		request,
		0,
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

func TestInitatePush_NoPush(t *testing.T) {
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
		&httptest.ResponseRecorder{},
	}

	writer := &responseWriter{
		testW,
		request,
		0,
	}

	writer.Header()["Link"] = []string{"</style.css>; rel=preload; nopush;"}

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
			"Foo": []string{"Baz"},
		},
	}
	testW := &testWriter{
		[]string{},
		nil,
		&httptest.ResponseRecorder{},
	}

	writer := &responseWriter{
		testW,
		request,
		0,
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

	if len(testW.options.Header["Foo"]) != 1 || testW.options.Header["Foo"][0] != "Baz" {
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
