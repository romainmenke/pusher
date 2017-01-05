package adaptive

import (
	"testing"
	"time"
)

func TestWeight(t *testing.T) {

	var (
		now            time.Time
		fiveMinutesAgo time.Time
		amount         float64
	)

	now = time.Now()
	fiveMinutesAgo = now.Add(time.Minute * -5)
	amount = 10

	amount = weight(amount, fiveMinutesAgo, now)

	if amount > 5.0001 {
		t.Fatal(amount)
	}
	if amount < 4.9999 {
		t.Fatal(amount)
	}
}

func BenchmarkWeight(b *testing.B) {

	var (
		now            time.Time
		fiveMinutesAgo time.Time
		amount         float64
	)

	now = time.Now()
	fiveMinutesAgo = now.Add(time.Minute * -5)
	amount = 10

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			amount = weight(amount, fiveMinutesAgo, now)
		}
	})
}
