package minimum_value_to_get_positive_step_by_step_sum

import "math"

func MinStartValue(nums []int) int {
	var minVal int
	var total int

	for i := 0; i < len(nums); i++ {
		total += nums[i]
		minVal = int(math.Min(float64(minVal), float64(total)))
	}
	return -minVal + 1
}
