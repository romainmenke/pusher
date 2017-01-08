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

RUNELOOP:
	for index, runeValue := range h {
		switch runeValue {
		case '<':
			linkStart = index + 1
		case '>':
			linkEnd = index
			break RUNELOOP
		case ';':
			linkStart = 0
			linkEnd = 0
		}
	}

	if linkStart == 0 || linkEnd == 0 {
		return ""
	}

	return strings.Replace(h[linkStart:linkEnd], " ", "", -1)
}
