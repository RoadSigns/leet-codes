package fizz_buzz

import "strconv"

func FizzBuzz(n int) []string {
	r := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		r = append(r, fizzOrBuzz(i))
	}
	return r
}

func fizzOrBuzz(i int) string {
	if i%5 == 0 && i%3 == 0 {
		return "FizzBuzz"
	}

	if i%5 == 0 {
		return "Buzz"
	}

	if i%3 == 0 {
		return "Fizz"
	}

	return strconv.Itoa(i)
}
