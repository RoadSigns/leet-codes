package minimized_string_length

func MinimizedStringLength(s string) int {
	l := make(map[rune]int)
	for _, i := range s {
		l[i]++
	}
	return len(l)
}
