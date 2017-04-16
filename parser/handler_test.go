package parser

import (
	"net/http"
	"testing"
	"time"

	"github.com/fd/httpmiddlewarevet"
)

func TestMiddleware(t *testing.T) {
	httpmiddlewarevet.Vet(t, func(h http.Handler) http.Handler {
		return Handler(h)
	})
}

func TestMiddlewareWithCache(t *testing.T) {
	httpmiddlewarevet.Vet(t, func(h http.Handler) http.Handler {
		return Handler(h, WithCache(), CacheDuration(time.Minute*5))
	})
}
