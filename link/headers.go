package link

import "net/http"

func copyPushSafeHeader(dst, src http.Header) {
	for k, vv := range src {
		switch k {
		case "Accept-Charset":
			dst[k] = vv
			continue
		case "Accept-Ch":
			dst[k] = vv
			continue
		case "Accept-Ch-Lifetime":
			dst[k] = vv
			continue
		case "Accept-Encoding":
			dst[k] = vv
			continue
		case "Accept-Language":
			dst[k] = vv
			continue
		case "Authorization":
			dst[k] = vv
			continue
		case "Cookie":
			dst[k] = vv
			continue
		case "Dnt":
			dst[k] = vv
			continue
		case "User-Agent":
			dst[k] = vv
			continue
		default:
			continue
		}
	}
}
