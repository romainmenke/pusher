[![Build Status](https://travis-ci.org/romainmenke/pusher.svg?branch=master)](https://travis-ci.org/romainmenke/pusher)
[![GoDoc](https://godoc.org/github.com/romainmenke/pusher?status.svg)](https://godoc.org/github.com/romainmenke/pusher)

Note : this requires golang 1.8 (will be released 31/01)

---

### What :

**pusher** is a collection of `http.HandlerFunc` to enable http2 Push Promises. Different strategies are implemented in different handlers.
The idea behind having multiple handlers is that H2 Push can also harm performance if the wrong strategy is used.

At the moment there is :
- adaptive : an experimental auto push handler.
- link : a push handler based on `Link` headers.
- linkheader : a static file based linkheader placer.

Planned :
- parser : a push handler that parses response data to generate push frames

Checkout the sub-packages for more details.

---

Response without Push :

![without push](https://raw.githubusercontent.com/romainmenke/pusher/master/example/readme/before_push.png)

Response with Push :

![with push](https://raw.githubusercontent.com/romainmenke/pusher/master/example/readme/after_push.png)

----
