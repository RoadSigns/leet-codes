package reverse_string

func ReverseString(s []byte) []byte {
	b := make([]byte, len(s))
	for i, c := range s {
		b[len(s)-i-1] = c
	}
	return b
}
