package average_salary_excluding_the_minimum_and_maximum_salary

import "math"

func Average(salary []int) float64 {
	max := 0
	min := math.MaxInt
	var total int

	for _, s := range salary {
		if min > s {
			min = s
		}

		if max < s {
			max = s
		}

		total += s
	}
	return float64(total-min-max) / float64(len(salary)-2)
}
