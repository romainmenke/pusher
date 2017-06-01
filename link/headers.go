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
	if vv, found := src["DNT"]; found {
		for _, v := range vv {
			dst.Add("DNT", v)
		}
	}
	if vv, found := src["Proxy-Authorization"]; found {
		for _, v := range vv {
			dst.Add("Proxy-Authorization", v)
		}
	}
	if vv, found := src["User-Agent"]; found {
		for _, v := range vv {
			dst.Add("User-Agent", v)
		}
	}
	// if vv, found := src["Host"]; found {
	// 	for _, v := range vv {
	// 		dst.Add("Host", v)
	// 	}
	// }
	// if vv, found := src["Max-Forwards"]; found {
	// 	for _, v := range vv {
	// 		dst.Add("Max-Forwards", v)
	// 	}
	// }
	// if vv, found := src["Origin"]; found {
	// 	for _, v := range vv {
	// 		dst.Add("Origin", v)
	// 	}
	// }
	// if vv, found := src["Referer"]; found {
	// 	for _, v := range vv {
	// 		dst.Add("Referer", v)
	// 	}
	// }
	if vv, found := src["TE"]; found {
		for _, v := range vv {
			dst.Add("TE", v)
		}
	}
	if vv, found := src["Transfer-Encoding"]; found {
		for _, v := range vv {
			dst.Add("Transfer-Encoding", v)
		}
	}
}
