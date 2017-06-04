package intsort

type Uints []uint

func (s Uints) Len() int {
	return len(s)
}

func (s Uints) Less(i int, j int) bool {
	return s[i] < s[j]
}

func (s Uints) Swap(i int, j int) {
	s[i], s[j] = s[j], s[i]
}
