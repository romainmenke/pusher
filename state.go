package pusher

import (
	"net/url"
	"sync"
	"time"
)

// pkg setup :
// create the state map
// create a mutex to protect the state map
func init() {

	mutex = &sync.Mutex{}
	pushMap = make(map[string]map[string]*push)

}

// mutex to protext the state map
var mutex *sync.Mutex

// the state map
var pushMap map[string]map[string]*push

// readFromPushMap is used to generate the paths for which Push Promises should be created.
func readFromPushMap(page string, writer func(path string)) {
	mutex.Lock()

	pagePush, found := pushMap[page]
	if found {
		trim(pagePush)
		for path, push := range pagePush {
			if push.weight > 2 {
				writer(path)
			}
		}
	}

	mutex.Unlock()
}

// addToPushMap is used to add a path to the state map
func addToPushMap(referer string, urlString string) {

	mutex.Lock()
	defer mutex.Unlock()

	ref := pathFromReferer(referer)
	if ref == "" {
		return
	}

	pa, found := pushMap[ref]
	if !found {
		pa = make(map[string]*push)
	}

	pu, found := pa[urlString]
	if !found {
		pu = &push{
			weight:     0,
			weightedAt: time.Now(),
		}
	}

	pu.weight++
	pa[urlString] = pu
	pushMap[ref] = pa

}

type push struct {
	weight     float64
	weightedAt time.Time
}

// pathFromReferer is wrapper around url.Parse that returns url.Path or an empty string
func pathFromReferer(str string) string {

	u, _ := url.Parse(str)
	if u == nil {
		return ""
	}

	return u.Path
}
