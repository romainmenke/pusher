package main

import (
	"encoding/json"
	"net/http"

	"github.com/romainmenke/pusher/casper"
	"github.com/romainmenke/pusher/link"
)

func main() {

	http.Handle("/",
		casper.Handler(
			link.Handler(
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

					// adding link headers is done manually in the example.
					// this better illustrates the workings of the push handler
					switch r.URL.Path {
					case "/":
						w.Header().Add("Link", "</css/stylesheet.css>; rel=preload; as=style;")
					case "/alpha.html":
						w.Header().Add("Link", "</css/stylesheet.css>; rel=preload; as=style;")
						w.Header().Add("Link", "</fonts/CutiveMono-Regular.ttf>; rel=preload; as=font;")
						w.Header().Add("Link", "</js/text_change.js>; rel=preload; as=script;")
					case "/beta.html":
						w.Header().Add("Link", "</css/stylesheet.css>; rel=preload; as=style;")
						w.Header().Add("Link", "</fonts/CutiveMono-Regular.ttf>; rel=preload; as=font;")
						w.Header().Add("Link", "</img/gopher.png>; rel=preload; as=image;")
						w.Header().Add("Link", "</img/gopher1.png>; rel=preload; as=image;")
						w.Header().Add("Link", "</img/gopher2.png>; rel=preload; as=image;")
						w.Header().Add("Link", "</img/gopher3.png>; rel=preload; as=image;")
						w.Header().Add("Link", "</img/gopher4.png>; rel=preload; as=image; nopush;")
						w.Header().Add("Link", "</img/gopher5.png>; rel=preload; as=image; nopush;")
					case "/gamma.html":
						w.Header().Add("Link", "</css/stylesheet.css>; rel=preload; as=style;")
						w.Header().Add("Link", "</fonts/CutiveMono-Regular.ttf>; rel=preload; as=font;")
						w.Header().Add("Link", "</call.json>; rel=preload; onlypush;") // onlypush not implemented yet
					default:
					}

					w.Header().Add("Cache-Control", "max-age=1200")

					http.FileServer(http.Dir("./example/static")).ServeHTTP(w, r)

				}),
			),
			casper.InferCookieMaxAgeFromResponse(),
			casper.DefaultCookieMaxAge(2400),
		),
	)

	http.HandleFunc("/call.json",
		apiCall,
	)

	err := http.ListenAndServeTLS(":4430", "./casper/example/localhost.crt", "./casper/example/localhost.key", nil)
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
