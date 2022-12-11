package reverse_string

func ReverseString(s []byte) []byte {
	b := make([]byte, len(s))
	for i, c := range s {
		b[len(s)-i-1] = c
	}
	return b
}

func ReverseStringByPointer(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
