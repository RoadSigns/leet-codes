package first_letter_to_appear_twice

func RepeatedCharacter(s string) byte {
	l := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, exists := l[s[i]]; exists {
			return s[i]
		}
		l[s[i]] = 1
	}

	return []byte(" ")[0]
}
