package sum

func SumOfUnique(nums []int) int {
	t := make(map[int]int)
	for _, n := range nums {
		t[n]++
	}

	total := 0
	for n, c := range t {
		if c == 1 {
			total += n
		}
	}

	return total
}
