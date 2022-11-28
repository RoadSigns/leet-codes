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
