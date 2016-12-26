WIP adaptive http2 Pusher

pusher will auto-magically generate Push Promises based on most served assets on a page by page basis.

The maths to determine which asset will be Pushed still need some fine tuning.

example :

```go
package main

import (
	"net/http"

	"github.com/romainmenke/pusher"
)

func main() {

	http.HandleFunc("/",
		pusher.Pusher(http.FileServer(http.Dir("./cmd/static")).ServeHTTP),
	)

	err := http.ListenAndServeTLS(":4430", "cmd/localhost.crt", "cmd/localhost.key", nil)
	if err != nil {
		panic(err)
	}

}
```

Response without Push :
![without push](https://raw.githubusercontent.com/romainmenke/pusher/master/cmd/readme/before_push.png)

Response with Push :
![with push](https://raw.githubusercontent.com/romainmenke/pusher/master/cmd/readme/after_push.png)
