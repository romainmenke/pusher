package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/romainmenke/pusher/parser"
)

func main() {

	http.Handle("/",
		parser.Handler(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.FileServer(http.Dir("./example/static")).ServeHTTP(w, r)
			}),
			parser.WithCache(),
			parser.CacheDuration(time.Hour*100000),
		),
	)

	http.HandleFunc("/call.json",
		apiCall,
	)

	go func() {
		err := http.ListenAndServe(":8000", nil)
		if err != nil {
			panic(err)
		}
	}()

	err := http.ListenAndServeTLS(":4430", "./link/example/localhost.crt", "./link/example/localhost.key", nil)
	if err != nil {
		panic(err)
	}

}

func apiCall(w http.ResponseWriter, r *http.Request) {
	a := struct {
		Some string
	}{Some: "Remote Data"}
	json.NewEncoder(w).Encode(a)
}
