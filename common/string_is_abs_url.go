package common

import "strings"

// IsAbsolute is a fast alternative for url.URL.IsAbs
// IsAbs benches at 300ns, IsAbsolute is good enough and does it in 12ns
// IsAbsolute will also focus on the start of any url string, so url length will not have a large impact on this.
func IsAbsolute(p string) bool {
	if strings.HasPrefix(p, "/") && !strings.HasPrefix(p, "//") {
		return false
	}

	return true
}
