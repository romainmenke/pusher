[![wercker status](https://app.wercker.com/status/e85096dae221207cf6685300fb9db8c3/s/master "wercker status")](https://app.wercker.com/project/byKey/e85096dae221207cf6685300fb9db8c3)
[![GoDoc](https://godoc.org/github.com/romainmenke/pusher?status.svg)](https://godoc.org/github.com/romainmenke/pusher)

Note : this requires golang 1.8 (will be released 31/01)

Note : wercker will fail until we update our CI flow to golang 1.8

---

### What :

**pusher** is an `http.HandlerFunc` to enable http2 Push Promises based on traffic.

### How :

**pusher** inspects the request headers to generate a mapping of pages and corresponding dependencies. Each dependency receives a weight based on how many times it is requested. This weight rapidly drops in amount.

### When is it great :

- you have a golang static file server
- you have a golang proxy server
- you receive a decent amount of traffic (±10 request / minute)

### How is it great :

Because **pusher** inspects the upstream headers and does not parse downstream data it is fast. Really fast. This makes it great for proxies. There is also no overhead for more pages, which is also great for proxies.

### Issues :

**benchmarks** show that the processing time scales almost linearly with the number of dependencies to be pushed for a certain page. This means that poorly designed pages get poor performance. The performance hit is still a lot less than the latency would be.

- 10 dependency reads : 800ns
- 100 dependency reads : 6000ns
- 1000 dependency reads : 62000ns

---

example :

```go
package main

import (
	"net/http"

	"github.com/romainmenke/pusher"
)

func main() {

	http.HandleFunc("/",
		pusher.HandlerFunc(http.FileServer(http.Dir("./cmd/static")).ServeHTTP),
	)

	err := http.ListenAndServeTLS(":4430", "cmd/localhost.crt", "cmd/localhost.key", nil)
	if err != nil {
		panic(err)
	}

}
```

---

A working example can be found in the cmd directory.

Setup :

- generate crt : `$ openssl req -x509 -sha256 -nodes -newkey rsa:2048 -days 365 -keyout localhost.key -out localhost.crt`
- place the key and crt into the cmd directoy.
- trust the crt in keychain.
- start with : `$ go run ./cmd/main.go`
- visit : https://localhost:4430/

---

Response without Push :

![without push](https://raw.githubusercontent.com/romainmenke/pusher/master/cmd/readme/before_push.png)

Response with Push :

![with push](https://raw.githubusercontent.com/romainmenke/pusher/master/cmd/readme/after_push.png)

----
