package link

import "strings"

const (
	NoPush  = "nopush"
	Preload = "rel=preload"
)

func parseLinkHeader(h string) string {

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

	if !strings.Contains(h[linkEnd+1:], Preload) {
		return ""
	}

	if strings.Contains(h[linkEnd+1:], NoPush) {
		return ""
	}

	res := strings.TrimSpace(h[linkStart:linkEnd])

	if isAbsolute(res) {
		return ""
	}

	return res
}
