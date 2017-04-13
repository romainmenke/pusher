package common

import "strings"

// ParseLinkHeader turns </fonts/CutiveMono-Regular.ttf>; rel=preload; as=font; into /fonts/CutiveMono-Regular.ttf
// It returns and empty string when:
// - the length limit is exceeded
// - the 'rel' directive does not have a value of 'preload'
// - the 'nopush' directive is present
func ParseLinkHeader(h string) string {

	var linkStart int
	var linkEnd int

	// Don't parse anything longer than the limit
	if len(h) > headerLengthLimit {
		return ""
	}

RUNELOOP:
	for index, runeValue := range h {
		switch runeValue {

		// This indicates the start of a url.
		case '<':
			linkStart = index + 1

			// This indicates the end of a url.
		case '>':
			linkEnd = index
			break RUNELOOP

			// This indicates the start of a directive.
			// Reset the indices
		case ';':
			linkStart = 0
			linkEnd = 0
		}

	}

	// If either one of the indices is 0, no url was found.
	if linkStart == 0 || linkEnd == 0 {
		return ""
	}

	// If the remainder does not contain the preload directive return an empty string.
	if !strings.Contains(h[linkEnd+1:], Preload) {
		return ""
	}

	// If the remainder contains the nopush directive return an empty string.
	if strings.Contains(h[linkEnd+1:], NoPush) {
		return ""
	}

	// Trim spaced from the url.
	url := strings.TrimSpace(h[linkStart:linkEnd])

	// Url must have one forward slash.
	if IsAbsolute(url) {
		return ""
	}

	return url
}
