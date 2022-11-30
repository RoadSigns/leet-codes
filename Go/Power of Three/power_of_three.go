package power_of_three

func IsPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}

	if n%3 == 0 {
		return IsPowerOfThree(n / 3)
	}

	if n == 1 {
		return true
	}

	return false
}
