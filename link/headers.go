package link

import "net/http"

func copyPushSafeHeader(dst, src http.Header) {

	if vv, found := src["Accept-Charset"]; found {
		for _, v := range vv {
			dst.Add("Accept-Charset", v)
		}
	}
	if vv, found := src["Accept-Encoding"]; found {
		for _, v := range vv {
			dst.Add("Accept-Encoding", v)
		}
	}
	if vv, found := src["Accept-Language"]; found {
		for _, v := range vv {
			dst.Add("Accept-Language", v)
		}
	}
	if vv, found := src["Authorization"]; found {
		for _, v := range vv {
			dst.Add("Authorization", v)
		}
	}
	if vv, found := src["Cookie"]; found {
		for _, v := range vv {
			dst.Add("Cookie", v)
		}
	}
	if vv, found := src["Dnt"]; found {
		for _, v := range vv {
			dst.Add("Dnt", v)
		}
	}
	if vv, found := src["User-Agent"]; found {
		for _, v := range vv {
			dst.Add("User-Agent", v)
		}
	}
}
