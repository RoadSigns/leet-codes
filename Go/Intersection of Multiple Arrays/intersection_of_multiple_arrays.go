package intersection_of_mulitple_arrays

import "sort"

func Intersection(nums [][]int) []int {
	a := make([]int, 0)
	l := make(map[int]int)
	k := len(nums)
	for _, n := range nums {
		for _, m := range n {
			l[m]++
			if l[m] == k {
				a = append(a, m)
			}
		}
	}
	sort.Ints(a)
	return a
}
