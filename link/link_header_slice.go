package link

func swap(s []string, i, j int) {
	s[i], s[j] = s[j], s[i]
}

func sortLinkHeaders(s []string) {
	var trueCounter int

	for i := range s {
		if parseLinkHeader(s[i]) != "" {
			swap(s, trueCounter, i)
			trueCounter++
		}
	}
}

func splitLinkHeaders(s []string) ([]string, []string) {
	var trueCounter int
	for i := range s {
		if parseLinkHeader(s[i]) != "" {
			swap(s, trueCounter, i)
			trueCounter++
		}
	}

	return s[:trueCounter], s[trueCounter:]
}

func splitLinkHeadersAndParse(s []string) ([]string, []string) {
	var trueCounter int
	var parsed string
	for i := range s {
		parsed = parseLinkHeader(s[i])
		if parsed != "" {
			s[i] = parsed
			swap(s, trueCounter, i)
			trueCounter++
		}
	}

	return s[:trueCounter], s[trueCounter:]
}
