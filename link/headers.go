package link

import "net/http"

func copyPushSafeHeader(dst, src http.Header) {
	if vv, found := src["Accept-Charset"]; found {
		dst["Accept-Charset"] = vv
	}
	if vv, found := src["Accept-Encoding"]; found {
		dst["Accept-Encoding"] = vv
	}
	if vv, found := src["Accept-Language"]; found {
		dst["Accept-Language"] = vv
	}
	if vv, found := src["Authorization"]; found {
		dst["Authorization"] = vv
	}
	if vv, found := src["Cookie"]; found {
		dst["Cookie"] = vv
	}
	if vv, found := src["Dnt"]; found {
		dst["Dnt"] = vv
	}
	if vv, found := src["User-Agent"]; found {
		dst["User-Agent"] = vv
	}
}
