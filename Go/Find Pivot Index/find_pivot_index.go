package find_pivot_index

func PivotIndex(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}

	leftTotal := 0
	for i, n := range nums {
		if leftTotal == (total - leftTotal - n) {
			return i
		}
		leftTotal += n
	}

	return -1
}
