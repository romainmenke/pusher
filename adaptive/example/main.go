package main

import (
	"encoding/json"
	"net/http"

	"github.com/romainmenke/pusher/adaptive"
)

func main() {

	http.HandleFunc("/",
		adaptive.Handler(http.FileServer(http.Dir("./cmd/static"))).ServeHTTP,
	)

	// json calls have been removed from pushed for now
	http.HandleFunc("/",
		adaptive.Handler(http.HandlerFunc(APICall)).ServeHTTP,
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
