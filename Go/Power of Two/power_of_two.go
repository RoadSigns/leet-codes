package power_of_two

func IsPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}

	if n == 1 {
		return true
	}

	if n%2 == 1 {
		return false
	}

	n = n / 2
	return IsPowerOfTwo(n)
}
