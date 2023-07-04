package first_unique_character_in_a_string

func FirstUniqChar(s string) int {
	l := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		l[s[i]]++
	}

	for i := 0; i < len(s); i++ {
		if l[s[i]] == 1 {
			return i
		}
	}

	return -1
}

func FirstUniqCharWithRuneValue(s string) int {
	counter := make([]int, 26)
	for _, c := range s {
		counter[c-'a']++
	}
	for i, c := range s {
		if counter[c-'a'] == 1 {
			return i
		}
	}
	return -1
}
