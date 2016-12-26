package main

import (
	"context"
	"net/http"

	"limbo.services/trace"
	"limbo.services/trace/dev"

	"github.com/romainmenke/pusher"
)

func main() {

	trace.DefaultHandler = dev.NewHandler(nil)

	http.HandleFunc("/",
		Tracer(
			pusher.Pusher(http.FileServer(http.Dir("./cmd/static")).ServeHTTP),
		),
	)

	err := http.ListenAndServeTLS(":4430", "cmd/localhost.crt", "cmd/localhost.key", nil)
	if err != nil {
		panic(err)
	}

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
