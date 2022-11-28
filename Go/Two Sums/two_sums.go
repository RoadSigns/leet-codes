package two_sums

func twoSum(nums []int, target int) []int {
	result := make([]int, 0, 2)
	for k, v := range nums {
		for i, j := range nums {
			if k == i {
				continue
			}
			if j+v == target {
				result = append(result, k, i)
				return result
			}
		}
	}
	return result
}

func twoSumTwo(nums []int, target int) []int {
	tmpMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if val, exists := tmpMap[target-nums[i]]; exists {
			return []int{val, i}
		} else {
			tmpMap[nums[i]] = i
		}
	}
	return nil
}
