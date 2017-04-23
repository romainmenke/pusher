package http2

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/romainmenke/pusher/link"
)

func TestServer_Push_Success_Link(t *testing.T) {

	errc := make(chan error, 3)
	wrapper := func(handler http.Handler) http.Handler {
		return link.Handler(http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				switch r.URL.RequestURI() {
				case "/":

					w.Header().Add("Link", "</pushed?get>; rel=preload")

				case "/pushed?get":

				default:
					errc <- fmt.Errorf("unknown RequestURL %q", r.URL.RequestURI())
				}

				handler.ServeHTTP(w, r)
			}))
	}

	test := Middleware_Push_Succes_Test_Factory(wrapper, errc)
	test(t)
}
