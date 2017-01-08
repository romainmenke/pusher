package link

import "strings"

func parseLinkHeader(h string) string {

	if !strings.Contains(h, "rel=preload") {
		return ""
	}
	if strings.Contains(h, "nopush") {
		return ""
	}

	var linkStart int
	var linkEnd int

	for index, runeValue := range h {
		switch runeValue {
		case '<':
			linkStart = index + 1
		case '>':
			linkEnd = index
		}
	}

	if linkStart == 0 || linkEnd == 0 {
		return ""
	}

	return h[linkStart:linkEnd]
}
