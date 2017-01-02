package main

import (
	"encoding/json"
	"net/http"

	"github.com/romainmenke/pusher"
)

func main() {

	http.HandleFunc("/",
		pusher.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Vary", "Accept-Encoding")
				w.Header().Set("Cache-Control", "public, max-age=7776000")
				http.FileServer(http.Dir("./cmd/static")).ServeHTTP(w, r)
			},
		),
	)

	// json calls have been removed from pushed for now
	http.HandleFunc("/call.json",
		pusher.HandlerFunc(APICall),
	)

	err := http.ListenAndServeTLS(":4430", "cmd/localhost.crt", "cmd/localhost.key", nil)
	if err != nil {
		panic(err)
	}

}

func APICall(w http.ResponseWriter, r *http.Request) {
	a := struct {
		Some string
	}{Some: "Remote Data"}
	json.NewEncoder(w).Encode(a)
}
