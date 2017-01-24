package main

import (
	"encoding/json"
	"net/http"

	"github.com/romainmenke/pusher/adaptive"
)

func main() {

	http.HandleFunc("/",
		adaptive.Handler(
			http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					http.FileServer(http.Dir("./example/static")).ServeHTTP(w, r)
				},
			),
		),
	)

	// json calls have been removed from pushed for now
	http.HandleFunc("/call.json",
		adaptive.Handler(http.HandlerFunc(APICall)),
	)

	err := http.ListenAndServeTLS(":4430", "./adaptive/example/localhost.crt", "./adaptive/example/localhost.key", nil)
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
