
### What :

**adaptive** is an `http.HandlerFunc` to enable http2 Push Promises based on traffic.

### How :

**adaptive** inspects the request headers to generate a mapping of pages and corresponding dependencies. Each dependency receives a weight based on how many times it is requested. This weight rapidly drops in amount.

### When is it great :

- you have a golang static file server
- you have a golang proxy server
- you receive a decent amount of traffic (±10 request / minute)
- you have cacheless clients

### How is it great :

Because **adaptive** inspects the request headers and does not parse response data it is fast. Really fast. This makes it great for proxies. There is also no overhead for more pages, which is also great for proxies.

### Issues :

Since **adapative** pushes those resources which always follow other resource it does not work when a client caches responses. It might still be useful for cacheless clients.

---

example :

```go
package main

import (
	"net/http"

	"github.com/romainmenke/pusher/adaptive"
)

func main() {

	http.HandleFunc("/",
		adaptive.HandlerFunc(http.FileServer(http.Dir("./cmd/static")).ServeHTTP),
	)

	err := http.ListenAndServeTLS(":4430", "cmd/localhost.crt", "cmd/localhost.key", nil)
	if err != nil {
		panic(err)
	}

}
```

---

A working example can be found in the example directory.

Setup :

- generate crt : `$ openssl req -x509 -sha256 -nodes -newkey rsa:2048 -days 365 -keyout localhost.key -out localhost.crt`
- place the key and crt into the cmd directoy.
- trust the crt in keychain.
- start with : `$ go run ./example/main.go`
- visit : https://localhost:4430/

---
