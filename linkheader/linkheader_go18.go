// +build go1.8

package linkheader

import (
	"net/http"
	"sync"

	"github.com/romainmenke/pusher/common"
)

func wrap(path string, assetMap map[string]struct{}, linkMap map[string][]string, m *sync.RWMutex, handler http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer handler.ServeHTTP(w, r)

		if r.Method != http.MethodGet && r.Method != http.MethodHead {
			return
		}

		m.RLock()
		defer m.RUnlock()

		if _, found := assetMap[r.URL.Path]; found {
			return
		}

		// Must be a pusher / Must not be behind a proxy / Must be proto 2 / Must be get
		if pusher, ok := w.(http.Pusher); ok && r.Header.Get(common.XForwardedFor) == "" && r.Method != http.MethodHead && r.ProtoMajor == 2 {

			for _, header := range linkMap[path] {
				pusher.Push(common.ParseLinkHeader(header), &http.PushOptions{Header: r.Header})
			}

		} else {

			for _, header := range linkMap[path] {
				w.Header().Add(common.Link, header)
			}

		}
	})
}
