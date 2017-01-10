package link

type LinkHeaderSlice []string

func (s LinkHeaderSlice) Len() int {
	return len(s)
}
func (s LinkHeaderSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s LinkHeaderSlice) Less(i, j int) bool {
	parsedI := parseLinkHeader(s[i])
	parsedJ := parseLinkHeader(s[j])

	if parsedI != "" && parsedJ == "" {
		return true
	}

	return false
}

func (s LinkHeaderSlice) Sort() {
	var trueCounter int

	for i := range s {
		if parseLinkHeader(s[i]) != "" {
			s.Swap(trueCounter, i) // swap
			trueCounter++          // next "true" item will move in the next slot
		}
	}
}

func (s LinkHeaderSlice) Split() ([]string, []string) {
	var trueCounter int

	for i := range s {
		if parseLinkHeader(s[i]) != "" {
			s.Swap(trueCounter, i) // swap
			trueCounter++          // next "true" item will move in the next slot
		}
	}

	return s[:trueCounter], s[trueCounter:]
}
