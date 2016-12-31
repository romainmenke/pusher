[![wercker status](https://app.wercker.com/status/e85096dae221207cf6685300fb9db8c3/s/master "wercker status")](https://app.wercker.com/project/byKey/e85096dae221207cf6685300fb9db8c3)
[![GoDoc](https://godoc.org/github.com/romainmenke/pusher?status.svg)](https://godoc.org/github.com/romainmenke/pusher)

Note : this requires golang 1.8

Note : wercker will fail until we update to golang 1.8

---

### What :

**pusher** is an `http.HandlerFunc` to enable http2 Push Promises based on traffic.

### How :

**pusher** inspects the request headers to generate a mapping of pages and corresponding dependencies. Each dependency receives a weight based on how many times it is requested. This weight rapidly drops in amount.

### Issues :

**pusher** can't push nested dependencies (e.g. fonts referenced in css files). At the moment I consider this an acceptable draw-back.

**benchmarks** indicate a scaling issue with large numbers of dependencies, the number of pages in state has no effect on performance. Will investigate.

After some work I got it down to these numbers :

- 10 dependency reads : 800ns
- 100 dependency reads : 6000ns
- 1000 dependency reads : 62000ns

If a single page has more than a 100 dependencies there are easier ways to optimize.

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
