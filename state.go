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
		trimmed := trim(pagePush)
		for path := range trimmed {
			writer(path)
		}
	}

	mutex.Unlock()
}

// addToPushMap is used to add a path to the state map
func addToPushMap(referer string, urlString string) {

	// lock state
	mutex.Lock()
	// unlock state
	defer mutex.Unlock()

	// get path
	ref := pathFromReferer(referer)
	if ref == "" {
		return
	}

	var (
		pagePushes map[string]*push
		found      bool
	)

	// // check if referer is a pushed asset
	// for parentPath, pagePushes2 := range pushMap {
	// 	for path := range pagePushes2 {
	// 		if path == ref {
	// 			ref = parentPath
	// 		}
	// 	}
	// }

	pagePushes, found = pushMap[ref]
	if !found {
		pagePushes = make(map[string]*push)
		pushMap[ref] = pagePushes
	}

	p, found := pagePushes[urlString]
	if !found {
		p = &push{
			weight:     0,
			weightedAt: time.Now(),
		}
		pagePushes[urlString] = p
	}

	p.weight++

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
