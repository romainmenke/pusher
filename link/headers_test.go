package link

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"testing"
)

var pushSaveHeaderDst http.Header

var lotsOfHeaders = http.Header{
	"Accept-Charset":  []string{"Accept-Charset-Value-A", "Accept-Charset-Value-B"},
	"Accept-Encoding": []string{"Accept-Encoding-Value-A", "Accept-Encoding-Value-B"},
	"Accept-Language": []string{"Accept-Language-Value-A", "Accept-Language-Value-B"},
	"Authorization":   []string{"Authorization-Value-A", "Authorization-Value-B"},
	"Cookie":          []string{"Cookie-Value-A", "Cookie-Value-B"},
	"Dnt":             []string{"Dnt-Value-A", "Dnt-Value-B"},
	"Proxy-Authorization":    []string{"Proxy-Authorization-Value-A", "Proxy-Authorization-Value-B"},
	"User-Agent":             []string{"User-Agent-Value-A", "User-Agent-Value-B"},
	"Host":                   []string{"Host-Value-A", "Host-Value-B"},
	"Max-Forwards":           []string{"Max-Forwards-Value-A", "Max-Forwards-Value-B"},
	"Origin":                 []string{"Origin-Value-A", "Origin-Value-B"},
	"Referer":                []string{"Referer-Value-A", "Referer-Value-B"},
	"TE":                     []string{"TE-Value-A", "TE-Value-B"},
	"Accept":                 []string{"Accept-Value-A", "Accept-Value-B"},
	"Accept-Datetime":        []string{"Accept-Datetime-Value-A", "Accept-Datetime-Value-B"},
	"Connection":             []string{"Connection-Value-A", "Connection-Value-B"},
	"Content-Length":         []string{"Content-Length-Value-A", "Content-Length-Value-B"},
	"Content-MD5":            []string{"Content-MD5-Value-A", "Content-MD5-Value-B"},
	"Content-Type":           []string{"Content-Type-Value-A", "Content-Type-Value-B"},
	"Date":                   []string{"Date-Value-A", "Date-Value-B"},
	"Expect":                 []string{"Expect-Value-A", "Expect-Value-B"},
	"Forwarded":              []string{"Forwarded-Value-A", "Forwarded-Value-B"},
	"From":                   []string{"From-Value-A", "From-Value-B"},
	"If-Match":               []string{"If-Match-Value-A", "If-Match-Value-B"},
	"If-Modified-Since":      []string{"If-Modified-Since-Value-A", "If-Modified-Since-Value-B"},
	"If-None-Match":          []string{"If-None-Match-Value-A", "If-None-Match-Value-B"},
	"If-Range":               []string{"If-Range-Value-A", "If-Range-Value-B"},
	"If-Unmodified-Since":    []string{"If-Unmodified-Since-Value-A", "If-Unmodified-Since-Value-B"},
	"Pragma":                 []string{"Pragma-Value-A", "Pragma-Value-B"},
	"Range":                  []string{"Range-Value-A", "Range-Value-B"},
	"Upgrade":                []string{"Upgrade-Value-A", "Upgrade-Value-B"},
	"Warning":                []string{"Warning-Value-A", "Warning-Value-B"},
	"Via":                    []string{"Via-Value-A", "Via-Value-B"},
	"Front-End-Https":        []string{"Front-End-Https-Value-A", "Front-End-Https-Value-B"},
	"Proxy-Connection":       []string{"Proxy-Connection-Value-A", "Proxy-Connection-Value-B"},
	"X-ATT-DeviceId":         []string{"X-ATT-DeviceId-Value-A", "X-ATT-DeviceId-Value-B"},
	"X-Correlation-ID":       []string{"X-Correlation-ID-Value-A", "X-Correlation-ID-Value-B"},
	"X-Csrf-Token":           []string{"X-Csrf-Token-Value-A", "X-Csrf-Token-Value-B"},
	"X-Forwarded-For":        []string{"X-Forwarded-For-Value-A", "X-Forwarded-For-Value-B"},
	"X-Forwarded-Host":       []string{"X-Forwarded-Host-Value-A", "X-Forwarded-Host-Value-B"},
	"X-Forwarded-Proto":      []string{"X-Forwarded-Proto-Value-A", "X-Forwarded-Proto-Value-B"},
	"X-Http-Method-Override": []string{"X-Http-Method-Override-Value-A", "X-Http-Method-Override-Value-B"},
	"X-Request-ID":           []string{"X-Request-ID-Value-A", "X-Request-ID-Value-B"},
	"X-Requested-With":       []string{"X-Requested-With-Value-A", "X-Requested-With-Value-B"},
	"X-UIDH":                 []string{"X-UIDH-Value-A", "X-UIDH-Value-B"},
	"X-Wap-Profile":          []string{"X-Wap-Profile-Value-A", "X-Wap-Profile-Value-B"},

	"Accept-Ch":          []string{"Accept-Ch-Value-A", "Accept-Ch-Value-B"},
	"Accept-Ch-Lifetime": []string{"Accept-Ch-Lifetime-Value-A", "Accept-Ch-Lifetime-Value-B"},
}

var fewHeaders = http.Header{
	"Accept-Charset":  []string{"Accept-Charset-Value-A", "Accept-Charset-Value-B"},
	"Accept-Encoding": []string{"Accept-Encoding-Value-A", "Accept-Encoding-Value-B"},
	"Accept-Language": []string{"Accept-Language-Value-A", "Accept-Language-Value-B"},
	"X-Request-ID":    []string{"X-Request-ID-Value-A", "X-Request-ID-Value-B"},
}

func BenchmarkPushSafeHeaders_LotsOfHeaders(b *testing.B) {

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		pushSaveHeaderDst = http.Header{}
		copyPushSafeHeader(pushSaveHeaderDst, lotsOfHeaders)

	}
}

func BenchmarkPushSafeHeaders_FewHeaders(b *testing.B) {

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		pushSaveHeaderDst = http.Header{}
		copyPushSafeHeader(pushSaveHeaderDst, fewHeaders)
	}
}

func pushSafeHeaderTest(t *testing.T, key string, header http.Header, safe bool) {

	if header.Get(key) == "" && safe {
		t.Fatal(key, "expected : value, got : nothing")
	}
	if header.Get(key) != "" && !safe {
		t.Fatal(key, "expected : empty, got : something", header.Get(key))
	}

	vv := header[key]

	if !safe && len(vv) == 0 {
		return
	}

	if len(vv) != 2 {
		t.Fatal(key, vv)
	}

	a := fmt.Sprintf("%s-Value-A", key)
	b := fmt.Sprintf("%s-Value-B", key)
	if vv[0] != a {
		t.Fatal(fmt.Sprintf("expected : %s, got : %s", a, vv[0]))
	}
	if vv[1] != b {
		t.Fatal(fmt.Sprintf("expected : %s, got : %s", b, vv[1]))
	}
}

func safeHeader(key string) bool {
	switch key {
	case "Accept-Charset":
	case "Accept-Ch":
	case "Accept-Ch-Lifetime":
	case "Accept-Encoding":
	case "Accept-Language":
	case "Authorization":
	case "Cookie":
	case "Dnt":
	case "User-Agent":
	default:
		return false

	}
	return true
}

func TestPushSafeHeaders_A(t *testing.T) {

	pushSaveHeaderDst = http.Header{}
	copyPushSafeHeader(pushSaveHeaderDst, lotsOfHeaders)

	for key := range pushSaveHeaderDst {
		pushSafeHeaderTest(t, key, pushSaveHeaderDst, safeHeader(key))
	}

}

var testHeaders = http.Header{
	"Accept-Charset":  []string{"Accept-Charset-Value-A", "Accept-Charset-Value-B"},
	"Accept-Encoding": []string{"Accept-Encoding-Value-A", "Accept-Encoding-Value-B"},
	"Dnt":             []string{"Dnt-Value-A", "Dnt-Value-B"},
	"Proxy-Authorization": []string{"Proxy-Authorization-Value-A", "Proxy-Authorization-Value-B"},
	"User-Agent":          []string{"User-Agent-Value-A", "User-Agent-Value-B"},
	"Host":                []string{"Host-Value-A", "Host-Value-B"},
}

func TestPushSafeHeaders_B(t *testing.T) {

	pushSaveHeaderDst = http.Header{}
	copyPushSafeHeader(pushSaveHeaderDst, testHeaders)

	for _, key := range []string{"Accept-Charset", "Accept-Encoding", "Dnt", "Proxy-Authorization", "User-Agent", "User-Agent", "Host"} {
		pushSafeHeaderTest(t, key, pushSaveHeaderDst, safeHeader(key))
	}

}

func TestAcceptCH(t *testing.T) {

	header := http.Header{}
	header.Set("Accept-CH", "foo")

	buf := bytes.NewBuffer(nil)
	err := header.Write(buf)
	if err != nil {
		t.Fatal(err)
	}

	strHeader := string(buf.Bytes())
	if !strings.Contains(strHeader, "Accept-Ch") {
		t.Fatal("expected \"Accept-Ch:...\"")
	}

}
