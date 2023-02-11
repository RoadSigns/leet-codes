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

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	ans := 0
	left := 0
	curr := 1

	for right := 0; right < len(nums); right++ {
		curr *= nums[right]
		for left <= right && curr >= k {
			curr /= nums[left]
			left++
		}
		ans += right - left + 1
	}

	return ans
}

func findBestSubarray(nums []int, k int) int {
	curr := 0
	for i := 0; i < k; i++ {
		curr += nums[i]
	}

	ans := curr
	for i := k; i < len(nums); i++ {
		curr += nums[i] - nums[i-k]
		ans = int(math.Max(float64(ans), float64(curr)))
	}
	return ans
}
