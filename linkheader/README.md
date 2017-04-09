
### What :

**linkheader** makes it easy to add `Link` headers for golang static file servers.

----

### How :

**link** reads from a static a file containing paths and corresponding headers. It uses a `http.ServeMux` to match routes.

### When is it great :

- you want to easily add `Link` headers.

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

				http.FileServer(http.Dir("./example/static")).ServeHTTP(w, r)

			}),
			linkheader.PathOption("./linkheader/example/linkheaders.txt"),
		),
	)

}
```

---

rules :

- start with the path you want to match
- add `Link` header values
- end with an empty line
- `-` is used to ignore a path. This allows you to match `/foo` but not `/foo/no-match`

```
/
</css/stylesheet.css>; rel=preload; as=style;

/foo
</css/stylesheet.css>; rel=preload; as=style;

/foo/no-match
-
```

---

note :

Links described in `Link` header values are ignored:

```
/
</css/stylesheet.css>; rel=preload; as=style;
```

`/css/stylesheet.css` would also match the `/` rule, but no `Link` headers will be set for this request.

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
