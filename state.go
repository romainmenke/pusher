package pusher

import (
	"net/http"
	"net/url"
	"sync"
	"time"
)

// pkg setup :
// create the state map
// create a mutex to protect the state map
func init() {

	mutex = &sync.RWMutex{}
	pushMap = make(map[string]map[string]*push)

}

// mutex to protect the state map
var mutex *sync.RWMutex

// the state map
var pushMap map[string]map[string]*push

// readFromPushMap is used to generate the paths for which Push Promises should be created.
func readFromPushMap(page string, writer func(path string)) {
	mutex.RLock()

	pagePush, found := pushMap[page]
	if found {
		trimmed := trim(pagePush)
		for path := range trimmed {
			writer(path)
		}
	}

	mutex.RUnlock()
}

// addToPushMap is used to add a path to the state map
func addToPushMap(request *http.Request) {

	// lock state
	mutex.Lock()
	// unlock state
	defer mutex.Unlock()

	// get initiator
	initiator := getInitiator(request)
	if initiator == "" {
		return
	}

	var (
		pagePushes map[string]*push
		found      bool
	)

	pagePushes, found = pushMap[initiator]
	if !found {
		pagePushes = make(map[string]*push)
		pushMap[initiator] = pagePushes
	}

	p, found := pagePushes[request.RequestURI]
	if !found {
		p = &push{
			weight:     0,
			weightedAt: time.Now(),
		}
		pagePushes[request.RequestURI] = p
	}

	p.weight++

}

type push struct {
	weight     float64
	weightedAt time.Time
}

func pathFromReferer(str string) string {
	u, _ := url.Parse(str)
	if u == nil {
		return ""
	}

	return u.Path
}
