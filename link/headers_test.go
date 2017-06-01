package link

import (
	"net/http"
	"testing"
)

var pushSaveHeaderDst http.Header

func BenchmarkPushSafeHeaders_LotsOfHeaders(b *testing.B) {

	src := http.Header{
		"Accept-Charset":  []string{"Some Header Value", "Other Header Value"},
		"Accept-Encoding": []string{"Some Header Value", "Other Header Value"},
		"Accept-Language": []string{"Some Header Value", "Other Header Value"},
		"Authorization":   []string{"Some Header Value", "Other Header Value"},
		"Cookie":          []string{"Some Header Value", "Other Header Value"},
		"DNT":             []string{"Some Header Value", "Other Header Value"},
		"Proxy-Authorization":    []string{"Some Header Value", "Other Header Value"},
		"User-Agent":             []string{"Some Header Value", "Other Header Value"},
		"Host":                   []string{"Some Header Value", "Other Header Value"},
		"Max-Forwards":           []string{"Some Header Value", "Other Header Value"},
		"Origin":                 []string{"Some Header Value", "Other Header Value"},
		"Referer":                []string{"Some Header Value", "Other Header Value"},
		"TE":                     []string{"Some Header Value", "Other Header Value"},
		"Accept":                 []string{"Some Header Value", "Other Header Value"},
		"Accept-Datetime":        []string{"Some Header Value", "Other Header Value"},
		"Connection":             []string{"Some Header Value", "Other Header Value"},
		"Content-Length":         []string{"Some Header Value", "Other Header Value"},
		"Content-MD5":            []string{"Some Header Value", "Other Header Value"},
		"Content-Type":           []string{"Some Header Value", "Other Header Value"},
		"Date":                   []string{"Some Header Value", "Other Header Value"},
		"Expect":                 []string{"Some Header Value", "Other Header Value"},
		"Forwarded":              []string{"Some Header Value", "Other Header Value"},
		"From":                   []string{"Some Header Value", "Other Header Value"},
		"If-Match":               []string{"Some Header Value", "Other Header Value"},
		"If-Modified-Since":      []string{"Some Header Value", "Other Header Value"},
		"If-None-Match":          []string{"Some Header Value", "Other Header Value"},
		"If-Range":               []string{"Some Header Value", "Other Header Value"},
		"If-Unmodified-Since":    []string{"Some Header Value", "Other Header Value"},
		"Pragma":                 []string{"Some Header Value", "Other Header Value"},
		"Range":                  []string{"Some Header Value", "Other Header Value"},
		"Upgrade":                []string{"Some Header Value", "Other Header Value"},
		"Warning":                []string{"Some Header Value", "Other Header Value"},
		"Via":                    []string{"Some Header Value", "Other Header Value"},
		"Front-End-Https":        []string{"Some Header Value", "Other Header Value"},
		"Proxy-Connection":       []string{"Some Header Value", "Other Header Value"},
		"X-ATT-DeviceId":         []string{"Some Header Value", "Other Header Value"},
		"X-Correlation-ID":       []string{"Some Header Value", "Other Header Value"},
		"X-Csrf-Token":           []string{"Some Header Value", "Other Header Value"},
		"X-Forwarded-For":        []string{"Some Header Value", "Other Header Value"},
		"X-Forwarded-Host":       []string{"Some Header Value", "Other Header Value"},
		"X-Forwarded-Proto":      []string{"Some Header Value", "Other Header Value"},
		"X-Http-Method-Override": []string{"Some Header Value", "Other Header Value"},
		"X-Request-ID":           []string{"Some Header Value", "Other Header Value"},
		"X-Requested-With":       []string{"Some Header Value", "Other Header Value"},
		"X-UIDH":                 []string{"Some Header Value", "Other Header Value"},
		"X-Wap-Profile":          []string{"Some Header Value", "Other Header Value"},
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		pushSaveHeaderDst = http.Header{}
		copyPushSafeHeader(pushSaveHeaderDst, src)

	}
}

func BenchmarkPushSafeHeaders_FewHeaders(b *testing.B) {

	src := http.Header{
		"Accept-Charset":  []string{"Some Header Value", "Other Header Value"},
		"Accept-Encoding": []string{"Some Header Value", "Other Header Value"},
		"Accept-Language": []string{"Some Header Value", "Other Header Value"},
		"X-Request-ID":    []string{"Some Header Value", "Other Header Value"},
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		pushSaveHeaderDst = http.Header{}
		copyPushSafeHeader(pushSaveHeaderDst, src)

	}
}
