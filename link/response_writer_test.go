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

func TestWrite(t *testing.T) {

	recorder := httptest.NewRecorder()
	writer := &responseWriter{
		recorder,
		nil,
		0,
		false,
	}

	writer.Write([]byte("foo"))
	if recorder.Body == nil {
		t.Fatal("nil body")
	}
	if "foo" != string(recorder.Body.Bytes()) {
		t.Fatal("expected foo, got :", recorder.Body.Bytes())
	}

	if writer.statusCode != 200 {
		t.Fatal()
	}
	if recorder.Code != 200 {
		t.Fatal()
	}
}

func TestWriteB(t *testing.T) {

	recorder := httptest.NewRecorder()
	writer := &responseWriter{
		recorder,
		nil,
		0,
		false,
	}

	writer.WriteHeader(301)
	writer.Write([]byte("foo"))
	if recorder.Body == nil {
		t.Fatal("nil body")
	}
	if "foo" != string(recorder.Body.Bytes()) {
		t.Fatal("expected foo, got :", recorder.Body.Bytes())
	}

	if writer.statusCode != 301 {
		t.Fatal()
	}
	if recorder.Code != 301 {
		t.Fatal()
	}
}

func TestWriteHeader(t *testing.T) {

	recorder := httptest.NewRecorder()
	writer := &responseWriter{
		recorder,
		nil,
		0,
		false,
	}

	writer.WriteHeader(301)

	if writer.statusCode != 301 {
		t.Fatal()
	}
	if recorder.Code != 301 {
		t.Fatal()
	}
}
