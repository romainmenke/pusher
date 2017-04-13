// +build go1.8

package parser

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestWrite(t *testing.T) {

	u, _ := url.Parse("/")

	request := &http.Request{
		Method: "GET",
		URL:    u,
	}

	recorder := httptest.NewRecorder()
	writer := newResponseWriter(recorder, request)

	writer.Write([]byte(testHTML))
	if recorder.Body == nil {
		t.Fatal("nil body")
	}
	if len(testHTML) != len(recorder.Body.Bytes()) {
		t.Fatal()
	}

	if 1024 != len(writer.body.Bytes()) {
		t.Fatal()
	}

	linkSlice := writer.ExtractLinks()

	t.Log(linkSlice)

}

func BenchmarkWrite(b *testing.B) {

	u, _ := url.Parse("/")

	request := &http.Request{
		Method: "GET",
		URL:    u,
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {

		recorder := httptest.NewRecorder()
		writer := newResponseWriter(recorder, request)

		writer.Write([]byte(testHTML))

		writer.ExtractLinks()

	}

}

var testHTML = `
<!DOCTYPE HTML>
<html>
<head>
	<link rel="stylesheet" type="text/css" href="/assets/css/gzip/bundle.min.css">
	<script type="text/javascript" src="/assets/js/gzip/bundle.min.js"></script>
</head>
</html>
`
