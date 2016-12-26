package pusher

import (
	"math"
	"net/http"
	"net/url"
	"sync"
	"time"
)

func init() {

	mutex = &sync.Mutex{}
	pushMap = make(map[string]*pagePush)

}

var mutex *sync.Mutex
var pushMap map[string]*pagePush

func Pusher(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer handler(w, r)

		if isPageGet(r) {
			if p, ok := w.(http.Pusher); ok {

				mutex.Lock()

				pagePush, found := pushMap[r.URL.String()]
				if found {
					trim(pagePush.pushes)
					for path, push := range pagePush.pushes {
						if push.weight > 2 {
							p.Push(path, nil)
						}
					}
				}

				mutex.Unlock()

			}
			return
		}

		go addToPushMap(r.Referer(), r.URL.String())
	})
}

func addToPushMap(referer string, urlString string) {

	mutex.Lock()
	defer mutex.Unlock()

	ref := pathFromReferer(referer)
	if ref == "" {
		return
	}

	pa, found := pushMap[ref]
	if !found {
		pa = &pagePush{
			pushes: make(map[string]*push),
		}
	}

	pu, found := pa.pushes[urlString]
	if !found {
		pu = &push{
			weight:     0,
			weightedAt: time.Now(),
		}
	}

	pu.weight++
	pa.pushes[urlString] = pu
	pushMap[ref] = pa

}

type pagePush struct {
	pushes map[string]*push
}

type push struct {
	weight     float64
	weightedAt time.Time
}

func pathFromReferer(str string) string {

	u, err := url.Parse(str)
	if err != nil {
		return ""
	}

	return u.Path
}

func trim(ps map[string]*push) {

	var max float64

	for _, p := range ps {
		weight(p)
		if p.weight > max {
			max = p.weight
		}
	}

	for key, p := range ps {
		if max*0.8 > p.weight {
			delete(ps, key)
		}
	}

}

func weight(p *push) {

	if p == nil {
		return
	}

	if p.weightedAt.IsZero() {
		p.weightedAt = time.Now()
		return
	}

	delta := time.Now().Sub(p.weightedAt)

	d := delta.Minutes() / 5

	value := decay(p.weight, 0.5, d)

	p.weight = value
	p.weightedAt = time.Now()
}

func decay(a, r, x float64) float64 {
	y := a * (math.Pow((1 - r), x))
	return y
}
