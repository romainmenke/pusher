package link

import "strings"

// isAbsolute is a fast alternative for url.URL.IsAbs
// IsAbs benches at 300ns, isAbsolute is good enough and does it in 12ns
// isAbsolute will also focus on the start of any url string, so url length will not have a large impact on this.
func isAbsolute(p string) bool {
	if strings.HasPrefix(p, "/") && !strings.HasPrefix(p, "//") {
		return false
	}

	return true
}
