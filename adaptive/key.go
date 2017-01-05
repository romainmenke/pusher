package adaptive

import (
	"net/http"
	"net/url"
	"strings"
)

func resourceKeyFromRequest(request *http.Request) string {

	if request == nil || request.URL == nil {
		return ""
	}

	return strings.TrimSuffix("https://"+request.Host+request.URL.Path, "/")

}

func resourceKeyFromURL(u *url.URL) string {

	if u == nil {
		return ""
	}

	return strings.TrimSuffix(u.Path, "/")

}

func ownerKeyFromRequest(request *http.Request) string {
	referer := request.Referer()
	pushReferer := request.Header.Get(PushRefererKey)
	if pushReferer != "" {
		referer = pushReferer
	}

	if referer == "" {
		return ""
	}

	u, _ := url.Parse(referer)
	if u == nil {
		return ""
	}

	return strings.TrimSuffix("https://"+u.Host+u.Path, "/")
}
