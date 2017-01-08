package link

import "net/url"

func isAbsolute(p string) bool {
	u, err := url.Parse(p)
	if err != nil {
		return false
	}

	return u.IsAbs()
}
