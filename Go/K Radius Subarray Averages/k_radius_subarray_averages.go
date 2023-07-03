package k_radius_subarray_averages

func GetAverages(nums []int, k int) []int {
	averages := make([]int, len(nums))

	for i := 0; i <= len(nums)-1; i++ {
		if i < k {
			averages[i] = -1
			continue
		}

		if i+k > len(nums)-1 {
			averages[i] = -1
			continue
		}

		averages[i] = getAverage(nums[i-k : i+k+1])
	}

	return averages
}

func getAverage(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total / len(nums)
}

func GetAverageUsingRadius(nums []int, k int) []int {
	radius := k*2 + 1
	result := make([]int, len(nums))
	left := 0
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]

		result[i] = -1
		if (radius-1)+left == i {
			result[k+left] = sum / radius
			sum -= nums[left]
			left += 1
		}
	}

	return result
}
