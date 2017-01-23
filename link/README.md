
### What :

**link** has two functions :

Determine H2 Push is possible.

`CanPush(w http.ResponseWriter, r *http.Request) bool`

Push every URL in the Link Header with `rel=preload` and without `nopush`

`Push(header http.Header, pusher http.Pusher)`

----

It is heavily based upon the cloudflare http2 Push implementation.

https://blog.cloudflare.com/announcing-support-for-http-2-server-push-2/

https://blog.cloudflare.com/http-2-server-push-with-multiple-assets-per-link-header/

### How :

**link** inspects the response headers to generate Push Promise frames.

### When is it great :

- you have a golang proxy server

### How is it great :

Because **link** inspects the response headers, it leaves Push config over to the content server which gives you a lot of flexibility and control over what gets pushed.

---

example :

```go
package main

import (
	"net/http"

	"github.com/romainmenke/pusher/link"
)

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
