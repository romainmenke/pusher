[![Build Status](https://travis-ci.org/romainmenke/pusher.svg?branch=master)](https://travis-ci.org/romainmenke/pusher)
[![MiddlewareVet](https://middleware.vet/github.com/romainmenke/pusher.svg)](https://middleware.vet#github.com/romainmenke/pusher)
[![Go Report Card](https://goreportcard.com/badge/github.com/romainmenke/pusher)](https://goreportcard.com/report/github.com/romainmenke/pusher)
[![codecov](https://codecov.io/gh/romainmenke/pusher/branch/master/graph/badge.svg)](https://codecov.io/gh/romainmenke/pusher)
[![GoDoc](https://godoc.org/github.com/romainmenke/pusher?status.svg)](https://godoc.org/github.com/romainmenke/pusher)


<p align="center">
  <img src="https://cloud.githubusercontent.com/assets/11521496/24838540/070645b2-1d4a-11e7-9c39-900371d5fda3.png" width="250"/>
</p>

# H2 Push Handlers

**pusher** is a collection of `http.Handler`'s to easily enable HTTP2 Push.

- [link](https://github.com/romainmenke/pusher/tree/master/link) : a H2 Push handler based on `Link` headers.
- [linkheader](https://github.com/romainmenke/pusher/tree/master/linkheader) : `Link` header placer.
- [parser](https://github.com/romainmenke/pusher/tree/master/parser) : html body parser -> generates Push Frames / Link Headers for you.

Checkout the sub-packages for more details.

----

You probably already saw this code snippet from the [go blog](https://blog.golang.org/h2push) :

```go
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if pusher, ok := w.(http.Pusher); ok {
            // Push is supported.
            if err := pusher.Push("/app.js", nil); err != nil {
                log.Printf("Failed to push: %v", err)
            }
        }
        // ...
    })
```

But obviously you don't want to hard code pushes for all your assets, especially in case of a proxy. That is where the [link](https://github.com/romainmenke/pusher/tree/master/link) package comes in. It reads the response headers and looks for `Link` headers. If found it transforms these into Pushes. This approach is based on how Cloudflare enables H2 Push.

Then you can simply do this:
```go
func main() {
	link.Handler(YourHandlerHere)
}
```

If you have a go server and don't have an easy method to add these `Link` headers you can checkout the [linkheader](https://github.com/romainmenke/pusher/tree/master/linkheader) package. It does all the heavy lifting for you.


----

The Go gopher was designed by Renee French. (http://reneefrench.blogspot.com/)
