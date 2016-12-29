package pusher

import (
	"fmt"
	"math/rand"
	"net/http"
	"testing"
)

func BenchmarkAddToPushMap(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			r := &http.Request{
				RequestURI: "/page-" + fmt.Sprint(rand.Intn(1000)),
				Header: http.Header{
					PushInitiatorKey: []string{"/site-" + fmt.Sprint(rand.Intn(10))},
				},
			}
			addToPushMap(r)
		}
	})
}

func BenchmarkReadFromPushMap(b *testing.B) {
	for n1 := 0; n1 < 100; n1++ {
		for n2 := 0; n2 < 5; n2++ {
			r := &http.Request{
				RequestURI: "/page-" + fmt.Sprint(n1),
				Header: http.Header{
					PushInitiatorKey: []string{"/site-" + fmt.Sprint(n2)},
				},
			}
			addToPushMap(r)
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
