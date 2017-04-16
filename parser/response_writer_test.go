package parser

import (
	"log"
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
	writer := getResponseWriter(recorder, request)
	defer writer.close()

	writer.Write([]byte(testHTML))
	if recorder.Body == nil {
		t.Fatal("nil body")
	}
	if len(testHTML) != len(recorder.Body.Bytes()) {
		t.Fatal()
	}

	if 1024 < len(writer.body.Bytes()) {
		t.Fatal()
	}

	linkSlice := writer.extractLinks()
	if len(linkSlice) != 3 {
		log.Println(linkSlice)
		t.Fatal(len(linkSlice), linkSlice)
	}

	found := 0
	for _, link := range linkSlice {
		switch link.Path() {
		case "/assets/css/gzip/bundle.min.css":
			found++
		case "/assets/js/gzip/bundle.min.js":
			found++
		case "/img":
			found++
		default:
			t.Fatal(link)
		}
	}

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
		writer := getResponseWriter(recorder, request)

		writer.Write([]byte(testHTML))

		writer.extractLinks()

		writer.close()

	}

}

var testHTML = `
<!DOCTYPE HTML>
<html>
<head>
	<!-- preload -->
	<link rel="preload" href="/assets/font.woff2" as="font" type="font/woff2">
	<link rel="preload" href="/style/other.css" as="style">
	<link rel="preload" href="//example.com/resource">
	<link rel="preload" href="https://fonts.example.com/font.woff2" as="font" crossorigin type="font/woff2">

	<!-- links -->
	<link rel="stylesheet" type="text/css" href="/assets/css/gzip/bundle.min.css">
	<link rel="stylesheet" type="text/css" href="foo.com/assets/css/gzip/bundle.min.css">
	<link rel="stylesheet" type="text/css" href="/">
	<script type="text/javascript" src="/assets/js/gzip/bundle.min.js"></script>
	<script type="text/javascript" src="foo.com/assets/js/gzip/bundle.min.js"></script>
	<script type="text/javascript" src="/"></script>
	<img src="/img" alt="some_text">
	<img src="foo.com/img" alt="some_text">
	<img src="/" alt="some_text">

	<!-- partial -->
	<link rel="stylesheet" type="text/css" href="/assets/css/gzip/partial`
