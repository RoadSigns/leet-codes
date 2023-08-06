package duplicates

func ContainsDuplicate(nums []int) bool {
	e := make(map[int]bool)
	for _, n := range nums {
		if e[n] {
			return true
		}
		e[n] = true
	}

	return false
}
