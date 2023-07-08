package ransom_note

func CanConstruct(ransomNote string, magazine string) bool {
	m := make(map[rune]int)
	for _, c := range magazine {
		m[c]++
	}

	for _, n := range ransomNote {
		if m[n] == 0 {
			return false
		}
		m[n]--
	}

	return true
}
