package link

import "strings"

// isAbsolute is a fast alternative for url.URL.IsAbs
// IsAbs benches at 300ns, this is good enough and does it in 13ns
// isAbsolute will also focus on the start of any url string, so url length will not have a large impact on this.
func isAbsolute(p string) bool {
	if strings.HasPrefix(p, "//") || strings.HasPrefix(p, "http://") || strings.HasPrefix(p, "https://") {
		return true
	}

	// if a '.' is found before '/' it is an absolute url
	// requesting "stylesheet.css" instead of "/stylesheet.css" is not ok and not supported
	for _, runeValue := range p {
		switch runeValue {
		case '.':
			return true
		case '/':
			return false
		}
	}

	return false
}
