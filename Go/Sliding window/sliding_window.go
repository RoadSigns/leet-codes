package sliding_window

import "math"

func findLength(nums []int, k int) int {
	left := 0
	curr := 0
	ans := 0.0

	for right := 0; right < len(nums); right++ {
		curr += nums[right]
		for curr > k {
			curr -= nums[left]
			left++
		}

		ans = math.Max(ans, float64(right-left+1))
	}

	return int(ans)
}
