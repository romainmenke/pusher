// +build go1.8

package casper

import (
	"errors"
	"net/http"
	"strings"
)

func (w *responseWriter) Push(target string, opts *http.PushOptions) error {

	pusher, ok := w.ResponseWriter.(http.Pusher)
	if !ok || pusher == nil {
		return http.ErrNotSupported
	}

	c := CasperFromContext(w.ctx)
	if c == nil {
		return errors.New("casper was not defined")
	}

	// Remove casper cookie header if it already exists.
	if cookies, ok := w.Header()["Set-Cookie"]; ok && len(cookies) != 0 {
		w.Header().Del("Set-Cookie")
		for _, cookieStr := range cookies {
			if strings.HasPrefix(cookieStr, defaultCookieName+"=") {
				continue
			}
			w.Header().Add("Set-Cookie", cookieStr)
		}
	}

	h := c.hash([]byte(target))

	// Check the content is already pushed or not.
	if search(w.hashValues, h) {
		return nil
	}

	if !c.skipPush {
		if err := pusher.Push(target, opts); err != nil {
			return err
		}
		w.somethingPushed = true
	}

	w.hashValues = append(w.hashValues, h)

	return nil
}

func (w *responseWriter) setCookie() error {
	c := CasperFromContext(w.ctx)
	if c == nil {
		return errors.New("casper was not defined")
	}

	if !w.somethingPushed {
		return nil
	}

	cookie, err := c.generateCookie(w.hashValues, w.Header())
	if err != nil {
		return err
	}
	http.SetCookie(w, cookie)

	return nil
}
