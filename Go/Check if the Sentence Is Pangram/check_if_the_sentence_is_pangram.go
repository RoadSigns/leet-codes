package check_if_the_sentence_is_pangram

func CheckIfPangram(sentence string) bool {
	l := make(map[byte]int)
	for i := 0; i < len(sentence); i++ {
		l[sentence[i]]++
	}
	return len(l) == 26
}
