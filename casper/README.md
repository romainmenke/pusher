[![Build Status](https://travis-ci.org/romainmenke/pusher.svg?branch=master)](https://travis-ci.org/romainmenke/pusher)
[![GoDoc](https://godoc.org/github.com/romainmenke/pusher?status.svg)](https://godoc.org/github.com/romainmenke/pusher/casper)

<p align="center">
  <img src="https://cloud.githubusercontent.com/assets/11521496/24838540/070645b2-1d4a-11e7-9c39-900371d5fda3.png" width="250"/>
</p>

# CASPer

I can't take any credit for the genius that makes this pkg work. I simply modified the awesome work of [tcnksm](https://github.com/tcnksm) on [go-casper](https://github.com/tcnksm/go-casper).

The original pkg did not work well as middleware. It was not possible to wrap the internals of the original as middleware without adding substantial overhead. This is why I decided to redesign it from the ground up as an `http.Handler` wrapper.


### How it works :

**casper** uses cookies to skip H2 Pushes.

It is based on [H2O](https://github.com/h2o/h2o)'s [CASPer](https://h2o.examp1e.net/configure/http2_directives.html#http2-casper) (cache-aware server-push).


### Why :

You don't hate returning visitors.

---

#### Stuff it does :

- read a cookie and determine what was pushed recently
- skip pushes that are thought to be in browser cache
- write a cookie which will be used on subsequent requests

It does not initiate any H2 Pushes.

----

### Note :

Until the tests in `casper_test.go` are implement this is definitely not production ready.

----

The Go gopher was designed by Renee French. (http://reneefrench.blogspot.com/)
