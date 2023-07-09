package check_if_all_characters_have_equal_number_of_occurances

func AreOccurrencesEqual(s string) bool {
	l := make(map[rune]int)
	for _, r := range s {
		l[r]++
	}

	count := make(map[int]int)
	for _, c := range l {
		count[c]++
		if len(count) > 1 {
			return false
		}
	}

	return true
}
