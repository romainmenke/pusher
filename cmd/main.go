package main

import (
	"net/http"

	"github.com/romainmenke/pusher"
)

func main() {

	http.HandleFunc("/", pusher.Pusher(Static()))

	err := http.ListenAndServeTLS(":4430", "cmd/localhost.crt", "cmd/localhost.key", nil)
	if err != nil {
		panic(err)
	}

}

func Static() func(w http.ResponseWriter, r *http.Request) {
	return http.FileServer(http.Dir("./cmd/static")).ServeHTTP
}
