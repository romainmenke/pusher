package linkheader

import (
	"net/http"
	"testing"

	"github.com/fd/httpmiddlewarevet"
)

func TestMiddleware(t *testing.T) {
	httpmiddlewarevet.Vet(t, func(h http.Handler) http.Handler {
		return Handler(h, PathOption("./linkheader/example/linkheaders.txt"))
	})
}
