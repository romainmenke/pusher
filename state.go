package pusher

import (
	"math"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// pkg setup :
// create the state map
// create a mutex to protect the state map
func init() {

	collectionMutex = &sync.RWMutex{}
	state = &collection{
		nodes: make(map[string]*node),
	}

}

// mutex to protect the state map
var collectionMutex *sync.RWMutex

// the state map
var state *collection

// readFromPushMap is used to generate the paths for which Push Promises should be created.
func readFromPushMap(path string, writer func(path string)) {

	// readFromPushMap_Old(path, writer)

	collectionMutex.RLock()
	defer collectionMutex.RUnlock()

	state.travel(path, writer)

}

// addToPushMap is used to add a path to the state map
func addToPushMap(request *http.Request, increment float64) {

	// addToPushMap_Old(request, increment)

	collectionMutex.Lock()
	defer collectionMutex.Unlock()

	// get initiator
	initiator := getInitiator(request)
	if initiator == "" {
		return
	}

	state.bind(request.RequestURI, initiator, increment)

}

func pathFromReferer(str string) string {
	u, _ := url.Parse(str)
	if u == nil {
		return ""
	}

	return u.Path
}

// State
// 1d collection of dependency pointers
// graph of mapped requests

type collection struct {
	nodes map[string]*node
}

func (c *collection) findOrCreate(resourcePath string) *node {

	var (
		found bool
		n     *node
	)

	n, found = c.nodes[resourcePath]
	if found {
		return n
	}

	n = &node{
		path:        resourcePath,
		connections: make(map[string]*connection),
	}

	c.nodes[resourcePath] = n
	return n
}

func (c *collection) find(resourcePath string) *node {

	var (
		found bool
		n     *node
	)

	n, found = c.nodes[resourcePath]
	if found {
		return n
	}

	return nil
}

func (c *collection) bind(resourcePath string, ownerPath string, amount float64) *collection {

	if ownerPath == "" {
		return c
	}

	var (
		resourceNode *node
		ownerNode    *node
		conn         *connection
		found        bool
	)

	resourceNode = c.findOrCreate(resourcePath)
	ownerNode = c.findOrCreate(ownerPath)

	conn, found = ownerNode.connections[resourcePath]
	if !found {
		conn = &connection{
			to:         resourceNode,
			weightedAt: time.Now(),
		}
		ownerNode.connections[resourcePath] = conn
	}

	conn.weight += amount

	return c

}

func (c *collection) travel(ownerPath string, pusher func(path string)) {

	var (
		max       float64
		now       time.Time
		ownerNode *node
	)

	now = time.Now()
	ownerNode = c.find(ownerPath)
	if ownerNode == nil {
		return
	}

	for resourcePath, conn := range ownerNode.connections {

		conn.weight = conn.weight * (math.Pow((0.5), (now.Sub(conn.weightedAt).Minutes() / 5)))
		conn.weightedAt = now

		if conn.weight < 0.8 {
			delete(ownerNode.connections, resourcePath)
		}
		if conn.weight > max {
			max = conn.weight
		}
	}

	max = max * 0.95

	for resourcePath, conn := range ownerNode.connections {
		if max < conn.weight && conn.weight > 10 {
			pusher(resourcePath)
		}

		if conn.to != nil && conn.to.connections != nil {
			c.travel(conn.to.path, pusher)
		}
	}
}

type node struct {
	path        string
	connections map[string]*connection
}

type connection struct {
	to         *node
	weight     float64
	weightedAt time.Time
}
