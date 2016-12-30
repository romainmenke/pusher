package pusher

import (
	"math"
	"time"
)

// trimmed applies weights to the state map and deletes paths below a threshold
// this screws up Scaling
// TODO: fix this
func trimmed(ps map[string]*push, writer func(path string)) {

	var max float64

	for key, p := range ps {
		weight(p)
		if p.weight < 0.8 {
			delete(ps, key)
		}
		if p.weight > max {
			max = p.weight
		}
	}

	for key, p := range ps {
		if max*0.8 < p.weight && p.weight > 2 {
			writer(key)
		}
	}
}

// weight applies a decay rate to paths.
// weight values will be halved every five minutes
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

// decay is the maths for the weight func
func decay(a, r, x float64) float64 {
	y := a * (math.Pow((1 - r), x))
	return y
}
