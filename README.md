[![Build Status](https://travis-ci.org/romainmenke/pusher.svg?branch=master)](https://travis-ci.org/romainmenke/pusher)
[![Go Report Card](https://goreportcard.com/badge/github.com/romainmenke/pusher)](https://goreportcard.com/report/github.com/romainmenke/pusher)
[![codecov](https://codecov.io/gh/romainmenke/pusher/branch/master/graph/badge.svg)](https://codecov.io/gh/romainmenke/pusher)
[![GoDoc](https://godoc.org/github.com/romainmenke/pusher?status.svg)](https://godoc.org/github.com/romainmenke/pusher)

<p align="center">
  <img src="https://cloud.githubusercontent.com/assets/11521496/24838540/070645b2-1d4a-11e7-9c39-900371d5fda3.png" width="250"/>
</p>

# H2 Push Handlers

**pusher** is a collection of `http.Handler`'s to easily enable HTTP2 Push.

- [link](https://github.com/romainmenke/pusher/tree/master/link) : a H2 Push handler based on `Link` headers.
- [casper](https://github.com/romainmenke/pusher/tree/master/casper) : CASPer handler.
- [rules](https://github.com/romainmenke/pusher/tree/master/rules) : Simple rules to generate `Link` headers or pushes.
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

But obviously you don't want to hard code pushes for all your assets, especially in case of a proxy. That is where these handlers come in. Just choose the right one for the job.

### Proxy Server

If you run a proxy server and want to enable H2 Push for all requests coming through, you implement the `link` pkg Handler and add `Link` headers on the source server. This approach is based on how Cloudflare enables H2 Push.

### Client Side Rendered

Client Side Rendered websites often have known critical assets like the js bundle. In this case it makes sense to have a couple of rules for which assets to push for a certain path. This is what the `rules` pkg does. It adds `Link` headers or sends Pushes depending on your setup based on simple rules.

### Server Side Rendered

A Server Side Rendered website with a CMS doesn't have known critical assets at deploy time. The `parser` pkg reads the first 1024 bytes from every html response and adds `Link` headers or sends Pushes depending on your setup.

----

The Go gopher was designed by Renee French. (http://reneefrench.blogspot.com/)
