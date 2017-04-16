package parser

import (
	"sync"
	"time"

	"github.com/romainmenke/pusher/common"
)

var (
	state     map[string]*cacheItem
	mutex     *sync.RWMutex
	cachePool *sync.Pool
	cacheTime time.Duration
)

type cacheItem struct {
	links   []common.Preloadable
	expires time.Time
}

func (c *cacheItem) reset() {
	c.links = nil
	c.expires = time.Time{}
}

func (c *cacheItem) close() {
	c.reset()
	cachePool.Put(c)
}

func init() {
	state = make(map[string]*cacheItem)
	mutex = &sync.RWMutex{}
	cachePool = &sync.Pool{
		New: func() interface{} {
			return &cacheItem{}
		},
	}
	cacheTime = time.Minute * 5
}

func getFromCache(path string) []common.Preloadable {
	mutex.RLock()
	defer mutex.RUnlock()

	item, found := state[path]
	if !found {
		return nil
	}

	if item.expires.Before(time.Now()) {
		state[path] = nil
		item.close()
		return nil
	}

	return item.links
}

func getCacheItemFromCache(path string) *cacheItem {
	mutex.RLock()
	defer mutex.RUnlock()

	item, found := state[path]
	if !found {
		return nil
	}

	if item.expires.Before(time.Now()) {
		state[path] = nil
		item.close()
		return nil
	}

	return item
}

func putOneInCache(path string, link common.Preloadable) {

	var item *cacheItem

	item = getCacheItemFromCache(path)

	if item == nil {
		item = cachePool.Get().(*cacheItem)
		item.reset()
		item.expires = time.Now().Add(cacheTime)
	}

	mutex.Lock()
	defer mutex.Unlock()

	item.links = append(item.links, link)
	state[path] = item

}
