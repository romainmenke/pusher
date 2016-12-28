package pusher

import (
	"net/http"
	"net/url"
	"testing"
)

func TestIsPageGet(t *testing.T) {

	u, _ := url.Parse("https://www.site.com/page?param=foo#section")
	r := &http.Request{
		URL:    u,
		Method: "GET",
	}

	if isPageGet(r) == false {
		t.Fatal()
	}

}

var benchBoolResult bool

func BenchmarkIsPageGet(b *testing.B) {
	u, _ := url.Parse("https://www.site.com/page?param=foo#section")
	r := &http.Request{
		URL:    u,
		Method: "GET",
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			benchBoolResult = isPageGet(r)
		}
	})
}
