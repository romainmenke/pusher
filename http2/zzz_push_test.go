package http2

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func Middleware_Push_Succes_Test_Factory(wrapper func(http.Handler) http.Handler, errc chan error) func(*testing.T) {

	test := func(t *testing.T) {
		const (
			mainBody = `<html>
index page
<link rel="stylesheet" type="text/css" href="/pushed?get">
</html>
`
			pushedBody = "<html>pushed page</html>"
			userAgent  = "testagent"
			cookie     = "testcookie"
		)

		var stURL string
		checkPromisedReq := func(r *http.Request, wantMethod string, wantH http.Header) error {
			if got, want := r.Method, wantMethod; got != want {
				return fmt.Errorf("promised Req.Method=%q, want %q", got, want)
			}
			if got, want := r.Header, wantH; !reflect.DeepEqual(got, want) {
				return fmt.Errorf("promised Req.Header=%q, want %q", got, want)
			}
			if got, want := "https://"+r.Host, stURL; got != want {
				return fmt.Errorf("promised Req.Host=%q, want %q", got, want)
			}
			if r.Body == nil {
				return fmt.Errorf("nil Body")
			}
			if buf, err := ioutil.ReadAll(r.Body); err != nil || len(buf) != 0 {
				return fmt.Errorf("ReadAll(Body)=%q,%v, want '',nil", buf, err)
			}
			return nil
		}

		st := newServerTester(t, wrapper(
			http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					switch r.URL.RequestURI() {
					case "/":

						w.Header().Set("Content-Type", "text/html")
						w.Header().Set("Content-Length", strconv.Itoa(len(mainBody)))
						w.WriteHeader(200)
						io.WriteString(w, mainBody)
						errc <- nil

					case "/pushed?get":
						wantH := http.Header{}
						if err := checkPromisedReq(r, "GET", wantH); err != nil {
							errc <- fmt.Errorf("/pushed?get: %v", err)
							return
						}
						w.Header().Set("Content-Type", "text/html")
						w.Header().Set("Content-Length", strconv.Itoa(len(pushedBody)))
						w.WriteHeader(200)
						io.WriteString(w, pushedBody)
						errc <- nil

					default:
						errc <- fmt.Errorf("unknown RequestURL %q", r.URL.RequestURI())
					}
				}),
		).ServeHTTP,
		)
		stURL = st.ts.URL

		// Send one request, which should push two responses.
		st.greet()
		getSlash(st)
		for k := 0; k < 2; k++ {
			select {
			case <-time.After(2 * time.Second):
				t.Errorf("timeout waiting for handler %d to finish", k)
			case err := <-errc:
				if err != nil {
					t.Fatal(err)
				}
			}
		}

		checkPushPromise := func(f Frame, promiseID uint32, wantH [][2]string) error {
			pp, ok := f.(*PushPromiseFrame)
			if !ok {
				return fmt.Errorf("got a %T; want *PushPromiseFrame", f)
			}
			if !pp.HeadersEnded() {
				return fmt.Errorf("want END_HEADERS flag in PushPromiseFrame")
			}
			if got, want := pp.PromiseID, promiseID; got != want {
				return fmt.Errorf("got PromiseID %v; want %v", got, want)
			}
			gotH := st.decodeHeader(pp.HeaderBlockFragment())
			if !reflect.DeepEqual(gotH, wantH) {
				return fmt.Errorf("got promised headers %v; want %v", gotH, wantH)
			}
			return nil
		}
		checkHeaders := func(f Frame, wantH [][2]string) error {
			hf, ok := f.(*HeadersFrame)
			if !ok {
				return fmt.Errorf("got a %T; want *HeadersFrame", f)
			}
			gotH := st.decodeHeader(hf.HeaderBlockFragment())

			for _, wantPair := range wantH {
				var found bool
				for _, gotPair := range gotH {
					if gotPair[0] == wantPair[0] && gotPair[1] == wantPair[1] {
						found = true
					}
				}
				if !found {
					return fmt.Errorf("got response headers %v; want %v", gotH, wantH)
				}
			}

			return nil
		}
		checkData := func(f Frame, wantData string) error {
			df, ok := f.(*DataFrame)
			if !ok {
				return fmt.Errorf("got a %T; want *DataFrame", f)
			}
			if gotData := string(df.Data()); gotData != wantData {
				return fmt.Errorf("got response data %q; want %q", gotData, wantData)
			}
			return nil
		}

		// Stream 1 has 2 PUSH_PROMISE + HEADERS + DATA
		// Stream 2 has HEADERS + DATA
		// Stream 4 has HEADERS
		expected := map[uint32][]func(Frame) error{
			1: {
				func(f Frame) error {
					return checkPushPromise(f, 2, [][2]string{
						{":method", "GET"},
						{":scheme", "https"},
						{":authority", st.ts.Listener.Addr().String()},
						{":path", "/pushed?get"},
					})
				},
				func(f Frame) error {
					return checkHeaders(f, [][2]string{
						{":status", "200"},
						{"content-type", "text/html"},
						{"content-length", strconv.Itoa(len(mainBody))},
					})
				},
				func(f Frame) error {
					return checkData(f, mainBody)
				},
			},
			2: {
				func(f Frame) error {
					return checkHeaders(f, [][2]string{
						{":status", "200"},
						{"content-type", "text/html"},
						{"content-length", strconv.Itoa(len(pushedBody))},
					})
				},
				func(f Frame) error {
					return checkData(f, pushedBody)
				},
			},
		}

		consumed := map[uint32]int{}
		for k := 0; len(expected) > 0; k++ {
			f, err := st.readFrame()
			if err != nil {
				for id, left := range expected {
					t.Errorf("stream %d: missing %d frames", id, len(left))
				}
				t.Fatalf("readFrame %d: %v", k, err)
			}
			id := f.Header().StreamID
			label := fmt.Sprintf("stream %d, frame %d", id, consumed[id])
			if len(expected[id]) == 0 {
				t.Fatalf("%s: unexpected frame %#+v", label, f)
			}
			check := expected[id][0]
			expected[id] = expected[id][1:]
			if len(expected[id]) == 0 {
				delete(expected, id)
			}
			if err := check(f); err != nil {
				t.Fatalf("%s: %v", label, err)
			}
			consumed[id]++
		}
	}

	return test
}
