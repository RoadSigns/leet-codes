package duplicates

import "sort"

func FindDuplicates(nums []int) []int {
	d := make(map[int]int)

	for _, n := range nums {
		d[n]++
	}

	result := make([]int, 0, len(d)/2)
	for i, c := range d {
		if c == 2 {
			result = append(result, i)
		}
	}
	sort.Ints(result)
	return result
}
