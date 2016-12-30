package pusher

import (
	"testing"
	"time"
)

func TestWeight(t *testing.T) {

	now := time.Now()
	fiveMinutesAgo := now.Add(time.Minute * -5)

	p := &push{
		weightedAt: fiveMinutesAgo,
		weight:     10,
	}

	weight(p, now)

	if p.weight > 5.0001 {
		t.Fatal(p.weight)
	}
	if p.weight < 4.9999 {
		t.Fatal(p.weight)
	}
}

func BenchmarkWeight(b *testing.B) {

	now := time.Now()
	fiveMinutesAgo := time.Now().Add(time.Minute * -5)

	p := &push{
		weightedAt: fiveMinutesAgo,
		weight:     10,
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			weight(p, now)
		}
	})
}
