package link

import "strings"

func isAbsolute(p string) bool {
	if strings.Contains(p, "http://") || strings.Contains(p, "https://") || strings.Contains(p, "www.") {
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
