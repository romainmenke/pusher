package http2

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/romainmenke/pusher/rules"
)

func TestServer_Push_Success_Rules(t *testing.T) {

	reader := strings.NewReader(`/
</pushed?get>; rel=preload;"

/pushed?get
-
`)

	errc := make(chan error, 3)
	wrapper := func(handler http.Handler) http.Handler {
		return rules.Handler(http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				switch r.URL.RequestURI() {
				case "/":

				case "/pushed?get":

				default:
					errc <- fmt.Errorf("unknown RequestURL %q", r.URL.RequestURI())
				}

				handler.ServeHTTP(w, r)
			}),
			rules.ReaderOption(reader),
		)
	}

	test := Middleware_Push_Succes_Test_Factory(wrapper, errc)
	test(t)
}
