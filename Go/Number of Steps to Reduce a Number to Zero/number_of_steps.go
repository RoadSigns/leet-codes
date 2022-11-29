package number_of_steps

func NumberOfSteps(num int) int {
	steps := 0
	for num > 0 {
		if num%2 == 0 {
			num /= 2
			steps++
			continue
		}

		num -= 1
		steps++
	}
	return steps
}
