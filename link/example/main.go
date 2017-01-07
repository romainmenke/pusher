package main

import (
	"encoding/json"
	"net/http"

	"github.com/romainmenke/pusher/link"
)

func main() {

	http.HandleFunc("/",
		link.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Vary", "Accept-Encoding")
				w.Header().Set("Cache-Control", "public, max-age=7776000")

				if r.URL.RequestURI() == "/" {

					w.Header().Set("Link", "</css/stylesheet.css>; rel=preload; as=style")
				}
				http.FileServer(http.Dir("./example/static")).ServeHTTP(w, r)
			},
		),
	)

	// json calls have been removed from pushed for now
	http.HandleFunc("/call.json",
		link.HandlerFunc(APICall),
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
