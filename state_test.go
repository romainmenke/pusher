package pusher

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkAddToPushMap(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			addToPushMap("/site-"+fmt.Sprint(rand.Intn(10)), "/page-"+fmt.Sprint(rand.Intn(1000)))
		}
	})
}

func BenchmarkReadFromPushMap(b *testing.B) {
	for n1 := 0; n1 < 100; n1++ {
		for n2 := 0; n2 < 5; n2++ {
			addToPushMap("/site-"+fmt.Sprint(n2), "/page-"+fmt.Sprint(n1))
		}
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			readFromPushMap("/site-"+fmt.Sprint(rand.Intn(5)), func(path string) {
				_ = path
			})
		}
	})
}
