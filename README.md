[![wercker status](https://app.wercker.com/status/e85096dae221207cf6685300fb9db8c3/s/master "wercker status")](https://app.wercker.com/project/byKey/e85096dae221207cf6685300fb9db8c3)
[![GoDoc](https://godoc.org/github.com/romainmenke/pusher?status.svg)](https://godoc.org/github.com/romainmenke/pusher)

Note : this requires golang 1.8

Note : wercker will fail until we update to golang 1.8

Note : can't push nested dependencies (e.g. fonts referenced in css files)

---

WIP adaptive http2 Pusher

pusher will auto-magically generate Push Promises based on most served assets on a page by page basis.

The maths to determine which asset will be Pushed still need some fine tuning.

**Unlike most auto-push middleware pusher does not parse any assets in search of urls. It simply inspects the request headers to generate a 1 level deep asset map for each page. This enables it to be useful single page apps and allows it to push API calls too.**

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
		pusher.Handler(http.FileServer(http.Dir("./cmd/static")).ServeHTTP),
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
