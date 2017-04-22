// +build !go1.8

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

		for _, header := range linkMap[path] {
			w.Header().Add(common.Link, header)
		}
	})
}
