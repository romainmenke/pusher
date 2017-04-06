
### What :

**linkheader** makes it easy to add `Link` headers for golang static file servers.

----

### How :

**link** reads from a static a file containing paths and corresponding headers. It uses a `http.ServeMux` to match routes.

### When is it great :

- you are lazy and want to easily add `Link` headers.

### Why :

For our single page web apps we use a golang static file server hosted on Heroku, which is not HTTP/2.0 capable. We use a proxy to upgrade to HTTP/2.0 and wanted to enable H2 Push.

---

example :

```go
package main

import (
	"encoding/json"
	"net/http"

	"github.com/romainmenke/pusher/linkheader"
)

func main() {

	http.Handle("/",
		linkheader.Handler(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

				w.Header().Set("Cache-Control", "public, max-age=86400")

				http.FileServer(http.Dir("./example/static")).ServeHTTP(w, r)
			}),
			"./linkheader/example/linkheaders.txt",
		),
	)

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
