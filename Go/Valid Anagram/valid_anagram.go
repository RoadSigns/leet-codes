package valid_anagram

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := make(map[rune]int, len(s))
	tMap := make(map[rune]int, len(t))

	for _, r := range s {
		sMap[r]++
	}

	for _, r := range t {
		tMap[r]++
	}

	if len(tMap) != len(sMap) {
		return false
	}

	for r, i := range sMap {
		if i != tMap[r] {
			return false
		}
	}

	return true
}
