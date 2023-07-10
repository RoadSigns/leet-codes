package largest_unique_number

func LargestUniqueNumber(nums []int) int {
	count := make(map[int]int)
	for _, n := range nums {
		count[n]++
	}

	largest := -1
	for k, v := range count {
		if v == 1 {
			if k > largest {
				largest = k
			}
		}
	}

	return largest
}
