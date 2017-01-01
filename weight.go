package pusher

import (
	"math"
	"time"
)

// trimmed applies weights to the state map and deletes paths below a threshold
// this screws up Scaling
// TODO: fix this
func trimmed(page map[string]*dependency, writer func(path string)) {

	var max float64
	now := time.Now()

	for key, d := range page {
		weight(d, now)
		if d.weight < 0.8 {
			delete(page, key)
		}
		if d.weight > max {
			max = d.weight
		}
	}

	for key, d := range page {
		if max*0.8 < d.weight && d.weight > 10 {
			writer(key)
		}

		if d.dependencies != nil {
			trimmed(d.dependencies, writer)
		}
	}
}

// weight applies a decay rate to paths.
// weight values will be halved every five minutes
func weight(d *dependency, now time.Time) {
	d.weight = d.weight * (math.Pow((0.5), (now.Sub(d.weightedAt).Minutes() / 5)))
	d.weightedAt = now
}

// decay is the maths for the weight func
func decay(a, r, x float64) float64 {
	y := a * (math.Pow((1 - r), x))
	return y
}
