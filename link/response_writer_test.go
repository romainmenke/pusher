package link

import (
	"net/http"
	"testing"

	httpmiddlewarevet "github.com/fd/httpmiddlewarevet"
)

func Test(t *testing.T) {
	httpmiddlewarevet.Vet(t, func(h http.Handler) http.Handler {
		return HandleFunc(h.ServeHTTP)
	})
}
