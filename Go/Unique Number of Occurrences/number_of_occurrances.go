package occurances

func UniqueOccurrences(arr []int) bool {
	n := make(map[int]int)
	for _, i := range arr {
		n[i]++
	}

	un := make(map[int]bool)
	for _, c := range n {
		if un[c] {
			return false
		}
		un[c] = true
	}

	return true
}
