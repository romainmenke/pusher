package pusher

import "net/http"

func ExamplePusher() {

	// Pusher wraps around the static file HandlerFunc
	http.HandleFunc("/",
		Pusher(http.FileServer(http.Dir("./cmd/static")).ServeHTTP),
	)

	err := http.ListenAndServeTLS(":4430", "cmd/localhost.crt", "cmd/localhost.key", nil)
	if err != nil {
		panic(err)
	}
}
