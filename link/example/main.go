package main

import (
	"encoding/json"
	"net/http"

	"github.com/romainmenke/pusher/link"
)

func main() {

	http.HandleFunc("/",
		HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {

				// adding link headers is done manually in the example.
				// this better illustrates the workings of the push handler
				switch r.URL.RequestURI() {
				case "/":
					w.Header().Add("Link", "</css/stylesheet.css>; rel=preload; as=style;")
					w.Header().Add("link", "</fonts/CutiveMono-Regular.ttf>; rel=preload; as=font;")
				case "/alpha.html":
					w.Header().Add("link", "</css/stylesheet.css>; rel=preload; as=style;")
					w.Header().Add("link", "</fonts/CutiveMono-Regular.ttf>; rel=preload; as=font;")
					w.Header().Add("link", "</js/text_change.js>; rel=preload; as=script;")
				case "/beta.html":
					w.Header().Add("link", "</css/stylesheet.css>; rel=preload; as=style;")
					w.Header().Add("link", "</fonts/CutiveMono-Regular.ttf>; rel=preload; as=font;")
					w.Header().Add("link", "</img/gopher.png>; rel=preload; as=image;")
					w.Header().Add("link", "</img/gopher1.png>; rel=preload; as=image;")
					w.Header().Add("link", "</img/gopher2.png>; rel=preload; as=image;")
					w.Header().Add("link", "</img/gopher3.png>; rel=preload; as=image;")
					w.Header().Add("link", "</img/gopher4.png>; rel=preload; as=image;")
					w.Header().Add("link", "</img/gopher5.png>; rel=preload; as=image;")
				case "/gamma.html":
					w.Header().Add("link", "</css/stylesheet.css>; rel=preload; as=style;")
					w.Header().Add("link", "</fonts/CutiveMono-Regular.ttf>; rel=preload; as=font;")
					w.Header().Add("link", "</call.json>; rel=preload;")
				default:
				}

				http.FileServer(http.Dir("./example/static")).ServeHTTP(w, r)
			},
		),
	)

	// json calls have been removed from pushed for now
	http.HandleFunc("/call.json",
		HandlerFunc(APICall),
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

func HandlerFunc(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return newPushHandlerFunc(handlerFunc)
}

func newPushHandlerFunc(handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if !link.CanPush(w, r) {
			handler(w, r)
			return
		}

		var p pusher
		p = pusher{writer: w}
		handler(&p, r)

	})
}

type pusher struct {
	writer http.ResponseWriter
}

func (p *pusher) Header() http.Header {
	return p.writer.Header()
}

func (p *pusher) Write(b []byte) (int, error) {
	return p.writer.Write(b)
}

func (p *pusher) WriteHeader(rc int) {
	link.Push(p.Header(), p.writer.(http.Pusher))
	p.writer.WriteHeader(rc)
}
