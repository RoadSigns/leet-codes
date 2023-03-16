package minimum_value_to_get_positive_step_by_step_sum

func MinStartValue(nums []int) int {
	current := 1
	for {
		tmp := current
		for _, i := range nums {
			tmp += i
			if tmp < 1 {
				current++
				break
			}
		}
		if tmp < 1 {
			continue
		}
		return current
	}
}
