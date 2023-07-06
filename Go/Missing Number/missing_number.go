package missing_number

func MissingNumber(nums []int) int {
	n := len(nums)
	s := make(map[int]int, n)
	for _, i := range nums {
		s[i]++
	}

	for i := 0; i < n; i++ {
		if _, exists := s[i]; exists {
			continue
		}
		return i
	}
	return -1
}
