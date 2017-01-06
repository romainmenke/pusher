
### What :

**link** is an `http.HandlerFunc` to enable http2 Push Promises based on downstream `Link` headers.

It is heavily based upon the cloudflare http2 Push implementation.

https://blog.cloudflare.com/announcing-support-for-http-2-server-push-2/

### How :

**link** inspects the response headers to generate Push Promise frames.

### When is it great :

- you have a golang proxy server

### How is it great :

Because **link** inspects the downstream headers it leaves Push config over to the content server which gives you a lot of flexibility and control over what gets pushed.

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
		link.HandlerFunc(http.FileServer(http.Dir("./cmd/static")).ServeHTTP),
	)

	err := http.ListenAndServeTLS(":4430", "cmd/localhost.crt", "cmd/localhost.key", nil)
	if err != nil {
		panic(err)
	}

}
```

---
