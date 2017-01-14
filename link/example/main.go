package main

import (
	"encoding/json"
	"net/http"

	"github.com/romainmenke/pusher/link"
)

func main() {

	http.HandleFunc("/",
		link.HandleFunc(
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
		link.HandleFunc(APICall),
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
