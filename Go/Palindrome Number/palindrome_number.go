package palindrome_number

import "strconv"

func IsPalindrome(x int) bool {
	runes := []rune(strconv.Itoa(x))
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	i, _ := strconv.Atoi(string(runes))
	return x == i
}

func IsPalindromeWithoutStringConversion(x int) bool {
	
}
