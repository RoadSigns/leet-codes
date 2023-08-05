package substring

import "math"

func LengthOfLongestSubstring(s string) int {
	chars := make(map[uint8]int)
	l := 0
	r := 0
	c := 0.0

	for r < len(s) {
		charR := s[r]
		chars[charR]++

		for chars[charR] > 1 {
			charL := s[l]
			chars[charL]--
			l++
		}

		c = math.Max(c, float64(r-l+1))
		r++
	}
	return int(c)
}
