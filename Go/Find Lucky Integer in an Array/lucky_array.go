package lucky_array

func FindLucky(arr []int) int {
	nums := make(map[int]int)
	for _, i := range arr {
		nums[i]++
	}

	l := -1
	for n, c := range nums {
		if n == c {
			if n > l {
				l = n
			}
		}
	}
	return l
}
