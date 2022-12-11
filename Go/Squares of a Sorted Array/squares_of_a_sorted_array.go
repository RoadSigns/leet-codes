package squares_of_a_sorted_array

func SquaresOfASortedArray(nums []int) []int {
	result := make([]int, len(nums))
	start := 0
	end := len(nums) - 1

	for i := len(nums) - 1; i >= 0; i-- {
		pS := pow(nums[start])
		pE := pow(nums[end])
		if pS > pE {
			result[i] = pS
			start++
		} else {
			result[i] = pE
			end--
		}
	}

	return result
}

func pow(n int) int {
	return n * n
}
