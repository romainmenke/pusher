package pusher

import (
	"fmt"
	"math/rand"
	"net/http"
	"testing"
)

// Test Scaling Of Pushes
func BenchmarkAddToPushMap1000x10(b *testing.B) {
	state = &collection{
		nodes: make(map[string]*node),
	}

	for n1 := 0; n1 < 1000; n1++ {
		for n2 := 0; n2 < 10; n2++ {
			r := &http.Request{
				RequestURI: "/page-" + fmt.Sprint(n2),
				Header: http.Header{
					PushRefererKey: []string{"/site-" + fmt.Sprint(n1)},
				},
			}
			addToPushMap(r, 1)
		}
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			r := &http.Request{
				RequestURI: "/page-" + fmt.Sprint(rand.Intn(1000)),
				Header: http.Header{
					PushRefererKey: []string{"/site-" + fmt.Sprint(rand.Intn(10))},
				},
			}
			addToPushMap(r, 1)
		}
	})
}

func BenchmarkReadFromPushMap1000x10(b *testing.B) {
	state = &collection{
		nodes: make(map[string]*node),
	}

	for n1 := 0; n1 < 1000; n1++ {
		for n2 := 0; n2 < 10; n2++ {
			r := &http.Request{
				RequestURI: "/page-" + fmt.Sprint(n2),
				Header: http.Header{
					PushRefererKey: []string{"/site-" + fmt.Sprint(n1)},
				},
			}
			addToPushMap(r, 1)
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

func BenchmarkAddToPushMap1000x100(b *testing.B) {
	state = &collection{
		nodes: make(map[string]*node),
	}

	for n1 := 0; n1 < 1000; n1++ {
		for n2 := 0; n2 < 100; n2++ {
			r := &http.Request{
				RequestURI: "/page-" + fmt.Sprint(n2),
				Header: http.Header{
					PushRefererKey: []string{"/site-" + fmt.Sprint(n1)},
				},
			}
			addToPushMap(r, 1)
		}
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			r := &http.Request{
				RequestURI: "/page-" + fmt.Sprint(rand.Intn(1000)),
				Header: http.Header{
					PushRefererKey: []string{"/site-" + fmt.Sprint(rand.Intn(10))},
				},
			}
			addToPushMap(r, 1)
		}
	})
}

func BenchmarkReadFromPushMap1000x100(b *testing.B) {
	state = &collection{
		nodes: make(map[string]*node),
	}

	for n1 := 0; n1 < 1000; n1++ {
		for n2 := 0; n2 < 100; n2++ {
			r := &http.Request{
				RequestURI: "/page-" + fmt.Sprint(n2),
				Header: http.Header{
					PushRefererKey: []string{"/site-" + fmt.Sprint(n1)},
				},
			}
			addToPushMap(r, 1)
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

func BenchmarkAddToPushMap1000x1000(b *testing.B) {
	state = &collection{
		nodes: make(map[string]*node),
	}

	for n1 := 0; n1 < 1000; n1++ {
		for n2 := 0; n2 < 1000; n2++ {
			r := &http.Request{
				RequestURI: "/page-" + fmt.Sprint(n2),
				Header: http.Header{
					PushRefererKey: []string{"/site-" + fmt.Sprint(n1)},
				},
			}
			addToPushMap(r, 1)
		}
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			r := &http.Request{
				RequestURI: "/page-" + fmt.Sprint(rand.Intn(1000)),
				Header: http.Header{
					PushRefererKey: []string{"/site-" + fmt.Sprint(rand.Intn(10))},
				},
			}
			addToPushMap(r, 1)
		}
	})
}

func BenchmarkReadFromPushMap1000x1000(b *testing.B) {
	state = &collection{
		nodes: make(map[string]*node),
	}

	for n1 := 0; n1 < 1000; n1++ {
		for n2 := 0; n2 < 1000; n2++ {
			r := &http.Request{
				RequestURI: "/page-" + fmt.Sprint(n2),
				Header: http.Header{
					PushRefererKey: []string{"/site-" + fmt.Sprint(n1)},
				},
			}
			addToPushMap(r, 1)
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

// Test Scaling Of Pages

func BenchmarkAddToPushMap10x10(b *testing.B) {
	state = &collection{
		nodes: make(map[string]*node),
	}

	for n1 := 0; n1 < 10; n1++ {
		for n2 := 0; n2 < 10; n2++ {
			r := &http.Request{
				RequestURI: "/page-" + fmt.Sprint(n2),
				Header: http.Header{
					PushRefererKey: []string{"/site-" + fmt.Sprint(n1)},
				},
			}
			addToPushMap(r, 1)
		}
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			r := &http.Request{
				RequestURI: "/page-" + fmt.Sprint(rand.Intn(1000)),
				Header: http.Header{
					PushRefererKey: []string{"/site-" + fmt.Sprint(rand.Intn(10))},
				},
			}
			addToPushMap(r, 1)
		}
	})
}

func BenchmarkReadFromPushMap10x10(b *testing.B) {
	state = &collection{
		nodes: make(map[string]*node),
	}

	for n1 := 0; n1 < 10; n1++ {
		for n2 := 0; n2 < 10; n2++ {
			r := &http.Request{
				RequestURI: "/page-" + fmt.Sprint(n2),
				Header: http.Header{
					PushRefererKey: []string{"/site-" + fmt.Sprint(n1)},
				},
			}
			addToPushMap(r, 1)
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

func BenchmarkAddToPushMap100x10(b *testing.B) {
	state = &collection{
		nodes: make(map[string]*node),
	}

	for n1 := 0; n1 < 100; n1++ {
		for n2 := 0; n2 < 10; n2++ {
			r := &http.Request{
				RequestURI: "/page-" + fmt.Sprint(n2),
				Header: http.Header{
					PushRefererKey: []string{"/site-" + fmt.Sprint(n1)},
				},
			}
			addToPushMap(r, 1)
		}
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			r := &http.Request{
				RequestURI: "/page-" + fmt.Sprint(rand.Intn(1000)),
				Header: http.Header{
					PushRefererKey: []string{"/site-" + fmt.Sprint(rand.Intn(10))},
				},
			}
			addToPushMap(r, 1)
		}
	})
}

func BenchmarkReadFromPushMap100x10(b *testing.B) {
	state = &collection{
		nodes: make(map[string]*node),
	}

	for n1 := 0; n1 < 100; n1++ {
		for n2 := 0; n2 < 10; n2++ {
			r := &http.Request{
				RequestURI: "/page-" + fmt.Sprint(n2),
				Header: http.Header{
					PushRefererKey: []string{"/site-" + fmt.Sprint(n1)},
				},
			}
			addToPushMap(r, 1)
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

func BenchmarkAddToPushMap1000x10b(b *testing.B) {
	state = &collection{
		nodes: make(map[string]*node),
	}

	for n1 := 0; n1 < 1000; n1++ {
		for n2 := 0; n2 < 10; n2++ {
			r := &http.Request{
				RequestURI: "/page-" + fmt.Sprint(n2),
				Header: http.Header{
					PushRefererKey: []string{"/site-" + fmt.Sprint(n1)},
				},
			}
			addToPushMap(r, 1)
		}
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			r := &http.Request{
				RequestURI: "/page-" + fmt.Sprint(rand.Intn(1000)),
				Header: http.Header{
					PushRefererKey: []string{"/site-" + fmt.Sprint(rand.Intn(10))},
				},
			}
			addToPushMap(r, 1)
		}
	})
}

func BenchmarkReadFromPushMap1000x10b(b *testing.B) {
	state = &collection{
		nodes: make(map[string]*node),
	}

	for n1 := 0; n1 < 1000; n1++ {
		for n2 := 0; n2 < 10; n2++ {
			r := &http.Request{
				RequestURI: "/page-" + fmt.Sprint(n2),
				Header: http.Header{
					PushRefererKey: []string{"/site-" + fmt.Sprint(n1)},
				},
			}
			addToPushMap(r, 1)
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
