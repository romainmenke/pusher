package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"limbo.services/trace"
	"limbo.services/trace/dev"

	"github.com/romainmenke/pusher"
)

func main() {

	trace.DefaultHandler = dev.NewHandler(nil)

	http.HandleFunc("/",
		Tracer(
			pusher.Handler(
				func(w http.ResponseWriter, r *http.Request) {
					w.Header().Set("Vary", "Accept-Encoding")
					w.Header().Set("Cache-Control", "public, max-age=7776000")
					fmt.Println(w.Header())
					http.FileServer(http.Dir("./cmd/static")).ServeHTTP(w, r)
				},
			),
		),
	)
	http.HandleFunc("/call.json",
		Tracer(
			pusher.Handler(APICall),
		),
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

func Tracer(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		span, _ := trace.New(context.Background(), "Tracer", trace.WithPanicGuard)
		defer span.Close()
		span.Metadata["request"] = r.URL.String()
		span.Metadata["referer"] = r.Referer()

		handler(w, r)
	})

}
