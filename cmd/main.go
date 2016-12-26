package main

import (
	"net/http"

	"github.com/romainmenke/pusher"
)

func main() {

	http.HandleFunc("/",
		pusher.Pusher(http.FileServer(http.Dir("./cmd/static")).ServeHTTP),
	)

	err := http.ListenAndServeTLS(":4430", "cmd/localhost.crt", "cmd/localhost.key", nil)
	if err != nil {
		panic(err)
	}

}
