[![Build Status](https://travis-ci.org/romainmenke/pusher.svg?branch=master)](https://travis-ci.org/romainmenke/pusher)
[![MiddlewareVet](https://middleware.vet/github.com/romainmenke/pusher.svg)](https://middleware.vet#github.com/romainmenke/pusher)
[![GoDoc](https://godoc.org/github.com/romainmenke/pusher?status.svg)](https://godoc.org/github.com/romainmenke/pusher/link)

<p align="center">
  <img src="https://cloud.githubusercontent.com/assets/11521496/24838540/070645b2-1d4a-11e7-9c39-900371d5fda3.png" width="250"/>
</p>

# Link

Heavily based upon the Cloudflare H2 Push implementation.

https://blog.cloudflare.com/announcing-support-for-http-2-server-push-2/

https://blog.cloudflare.com/http-2-server-push-with-multiple-assets-per-link-header/

### How it works :

**link** inspects response headers to generate Push Promises.

### Why :

You like speed.

---

#### Stuff it does :

- generates H2 Push frames based on `Link` headers
- respects `nopush`
- prevents recursive pushes
- is compatible with HTTP1.1 requests
- plays nice when behind a proxy (No Pushes when `X-Forwarded-For` is set)

---

example :

```go
package main

import (
	"net/http"

	"github.com/romainmenke/pusher/link"
)

func main() {

		link.Handler(YourHandlerHere)

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
