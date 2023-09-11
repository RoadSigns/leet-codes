package daily_temperatures

func DailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0)
	ans := make([]int, len(temperatures))

	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[j] = i - j
		}

		stack = append(stack, i)
	}

	return ans
}
