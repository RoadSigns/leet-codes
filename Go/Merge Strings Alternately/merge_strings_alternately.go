package merge_strings_alternately

func MergeAlternately(word1 string, word2 string) string {
	w1 := len(word1)
	w2 := len(word2)
	w := make([]byte, 0)

	l := 0

	if w1 < w2 {
		l = w2
	} else {
		l = w1
	}

	for i := 0; i < l; i++ {
		if i < w1 {
			w = append(w, word1[i])
		}

		if i < w2 {
			w = append(w, word2[i])
		}
	}

	return string(w)
}
