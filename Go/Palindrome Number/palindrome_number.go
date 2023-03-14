package palindrome_number

import (
	"strconv"
)

func IsPalindrome(x int) bool {
	runes := []rune(strconv.Itoa(x))
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	i, _ := strconv.Atoi(string(runes))
	return x == i
}

func IsPalindromeWithoutStringConversion(x int) bool {
	if x < 0 {
		return false
	}
	if x < 9 {
		return true
	}
	xC := x
	tmp := 0
	for x > 0 {
		tmp = (tmp * 10) + (x % 10)
		x = x / 10
	}
	return xC == tmp
}
