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
	pages = make(map[string]map[string]*dependency)
	dependencies = make(map[string]*dependency)

}

// mutex to protect the state map
var mutex *sync.RWMutex

// the state map
var pages map[string]map[string]*dependency
var dependencies map[string]*dependency

// readFromPushMap is used to generate the paths for which Push Promises should be created.
func readFromPushMap(path string, writer func(path string)) {
	mutex.RLock()

	page, found := pages[path]
	if found {
		trimmed(page, writer)
	}

	mutex.RUnlock()
}

// addToPushMap is used to add a path to the state map
func addToPushMap(request *http.Request, increment float64) {

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
		page  map[string]*dependency
		d     *dependency
		d1    *dependency
		found bool
	)

	d1, found = dependencies[initiator]
	if found {

		if d1.dependencies == nil {
			d1.dependencies = make(map[string]*dependency)
		}

		page = d1.dependencies

	} else {

		page, found = pages[initiator]
		if !found {
			page = make(map[string]*dependency)
			pages[initiator] = page
		}

	}

	d, found = dependencies[request.RequestURI]
	if !found {
		d = &dependency{
			weight:     0,
			weightedAt: time.Now(),
		}
	}

	dependencies[request.RequestURI] = d
	page[request.RequestURI] = d

	d.weight += increment

}

type dependency struct {
	weight       float64
	weightedAt   time.Time
	dependencies map[string]*dependency
}

func pathFromReferer(str string) string {
	u, _ := url.Parse(str)
	if u == nil {
		return ""
	}

	return u.Path
}
