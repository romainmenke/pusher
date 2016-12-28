package pusher

import (
	"testing"
	"time"
)

func TestWeight(t *testing.T) {

	fiveMinutesAgo := time.Now().Add(time.Minute * -5)

	p := &push{
		weightedAt: fiveMinutesAgo,
		weight:     10,
	}

	weight(p)

	if p.weight > 6 {
		t.Fatal(p.weight)
	}
	if p.weight < 4 {
		t.Fatal(p.weight)
	}
}

func BenchmarkWeight(b *testing.B) {

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			fiveMinutesAgo := time.Now().Add(time.Minute * -5)

			p := &push{
				weightedAt: fiveMinutesAgo,
				weight:     10,
			}

			weight(p)
		}
	})
}
