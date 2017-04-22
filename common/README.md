[![Build Status](https://travis-ci.org/romainmenke/pusher.svg?branch=master)](https://travis-ci.org/romainmenke/pusher)
[![MiddlewareVet](https://middleware.vet/github.com/romainmenke/pusher.svg)](https://middleware.vet#github.com/romainmenke/pusher)
[![GoDoc](https://godoc.org/github.com/romainmenke/pusher?status.svg)](https://godoc.org/github.com/romainmenke/pusher/common)

<p align="center">
  <img src="https://cloud.githubusercontent.com/assets/11521496/24838540/070645b2-1d4a-11e7-9c39-900371d5fda3.png" width="250"/>
</p>

# Common

### Preload Link Header

examples :

```
Link: <https://example.com/font.woff2>; rel=preload; as=font; type='font/woff2'
Link: <https://example.com/app/script.js>; rel=preload; as=script
Link: <https://example.com/logo-hires.jpg>; rel=preload; as=image
Link: <https://fonts.example.com/font.woff2>; rel=preload; as=font; crossorigin; type='font/woff2'
```
[source](https://w3c.github.io/preload/M)

---

### Preload HTML Link

```html
<link rel="preload" href="/assets/font.woff2" as="font" type="font/woff2">
<link rel="preload" href="/style/other.css" as="style">
<link rel="preload" href="//example.com/resource">
<link rel="preload" href="https://fonts.example.com/font.woff2" as="font" crossorigin type="font/woff2">
```

[source](https://w3c.github.io/preload/M)
