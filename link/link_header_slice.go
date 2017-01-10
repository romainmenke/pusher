package link

type ByPushable []string

func (s ByPushable) Len() int {
	return len(s)
}
func (s ByPushable) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByPushable) Less(i, j int) bool {
	parsedI := parseLinkHeader(s[i])
	parsedJ := parseLinkHeader(s[j])

	if parsedI != "" && parsedJ == "" {
		return true
	}

	return false
}

func (s ByPushable) Sort() {
	var trueCounter int

	for i := range s {
		if parseLinkHeader(s[i]) != "" {
			s.Swap(trueCounter, i)
			trueCounter++
		}
	}
}

func (s ByPushable) Split() ([]string, []string) {
	var trueCounter int

	for i := range s {
		if parseLinkHeader(s[i]) != "" {
			s.Swap(trueCounter, i)
			trueCounter++
		}
	}

	return s[:trueCounter], s[trueCounter:]
}
