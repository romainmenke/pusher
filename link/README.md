[![Build Status](https://travis-ci.org/romainmenke/pusher.svg?branch=master)](https://travis-ci.org/romainmenke/pusher)
[![MiddlewareVet](https://middleware.vet/github.com/romainmenke/pusher.svg)](https://middleware.vet#github.com/romainmenke/pusher)
[![GoDoc](https://godoc.org/github.com/romainmenke/pusher?status.svg)](https://godoc.org/github.com/romainmenke/pusher)

---

### Link :

It is heavily based upon the cloudflare H2 Push implementation.

https://blog.cloudflare.com/announcing-support-for-http-2-server-push-2/

https://blog.cloudflare.com/http-2-server-push-with-multiple-assets-per-link-header/

### How :

**link** inspects the response headers to generate Push Promise frames.

### Why :

You like speed.

---

#### Stuff it does :

- reads `Link` header values
- generates H2 Push frames
- respects `nopush`
- prevents recursive pushes
- is compatible with HTTP1.1 requests
- copies request headers to generated Push requests

---

example :

```go
package main

import (
	"net/http"

	"github.com/romainmenke/pusher/link"
)

func main() {

	http.HandleFunc("/",
		link.Handler(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

				// adding link headers is done manually in the example.
				// this better illustrates the workings of the Handler
				switch r.URL.RequestURI() {
				case "/":
					w.Header().Add("Link", "</css/stylesheet.css>; rel=preload; as=style;")
					w.Header().Add("Link", "</fonts/CutiveMono-Regular.ttf>; rel=preload; as=font;")
				default:
				}

				http.FileServer(http.Dir("./example/static")).ServeHTTP(w, r)
			}),
		).ServeHTTP,
	)

	err := http.ListenAndServeTLS(":4430", "./link/example/localhost.crt", "./link/example/localhost.key", nil)
	if err != nil {
		panic(err)
	}

}
```

---

Reference for `Link` headers :

https://w3c.github.io/preload/

| consumer | Preload directive |
|----------|-------------------|
| `<audio>, <video>` | `<link rel=preload as=media href=...>` |
| `<script>, Worker's importScripts` | `<link rel=preload as=script href=...>` |
| `<link rel=stylesheet>, CSS @import` | `<link rel=preload as=style href=...>` |
| `CSS @font-face` | `<link rel=preload as=font href=...>` |
| `<img>, <picture>, srcset, imageset` | `<link rel=preload as=image href=...>` |
| `SVG's <image>, CSS *-image` | `<link rel=preload as=image href=...>` |
| `XHR, fetch` | `<link rel=preload href=...>` |
| `Worker, SharedWorker` | `<link rel=preload as=worker href=...>` |
| `<embed>` | `<link rel=preload as=embed href=...>` |
| `<object>` | `<link rel=preload as=object href=...>` |
| `<iframe>, <frame>` | `<link rel=preload as=document href=...>` |

---
