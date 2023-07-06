package counting_elements

func CountElements(arr []int) int {
	t := 0
	n := make(map[int]int)
	for _, i := range arr {
		n[i]++
	}

	for _, i := range arr {
		if _, exists := n[i+1]; exists {
			t++
		}
	}
	return t
}
