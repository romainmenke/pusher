[![Build Status](https://travis-ci.org/romainmenke/pusher.svg?branch=master)](https://travis-ci.org/romainmenke/pusher)
[![MiddlewareVet](https://middleware.vet/github.com/romainmenke/pusher.svg)](https://middleware.vet#github.com/romainmenke/pusher/parser)
[![GoDoc](https://godoc.org/github.com/romainmenke/pusher?status.svg)](https://godoc.org/github.com/romainmenke/pusher/parser)

<p align="center">
  <img src="https://cloud.githubusercontent.com/assets/11521496/24838540/070645b2-1d4a-11e7-9c39-900371d5fda3.png" width="250"/>
</p>

# Parser

Auto generate Push Frames / Link Headers from html response bodies

----

Note :

Not ready for production

- has not been optimized
- few tests have been written
- might break everything. might be awesome

----

### How it works :

**parser** reads the first 1024 bytes from html response bodies. For H2 connections and Push capable clients it will generate Push Frames. For H1 connections or Push incapable clients it will send Link Headers. If a preload link was also found in the html the asset is ignored. Only one preload method should be used and this pkg should not alter the response body.

---

#### Stuff it does :

- generates H2 Push frames / Link Headers based on html bodies
- has a cache option (to only read from the body once)
- respects `nopush`
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

	http.Handle("/",
		parser.Handler(
				YourHandlerHere
			),
			parser.WithCache(),
			parser.CacheDuration(time.Hour*100000),
		),
	)

}
```

---

The Go gopher was designed by Renee French. (http://reneefrench.blogspot.com/)
