package maximum_average_subarray_i

import "math"

func findMaxAverage(nums []int, k int) float64 {
	curr := 0
	for i := 0; i < k; i++ {
		curr += nums[i]
	}

	ans := float64(curr)
	for i := k; i < len(nums); i++ {
		curr += nums[i] - nums[i-k]
		ans = math.Max(float64(ans), float64(curr))
	}

	return ans / float64(k)
}
