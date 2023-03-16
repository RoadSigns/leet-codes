package running_sum_of_1d_array

func RunningSum(nums []int) []int {
	result := make([]int, 0, len(nums))
	previous := 0
	for _, i := range nums {
		previous += i
		result = append(result, previous)
	}
	return result
}
