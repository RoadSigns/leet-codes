package maximum_number_of_balloons

func MaxNumberOfBalloons(text string) int {
	letters := make(map[rune]int)
	letters['b'] = 0
	letters['a'] = 0
	letters['l'] = 0
	letters['o'] = 0
	letters['n'] = 0

	for _, l := range text {
		if _, exists := letters[l]; exists {
			letters[l]++
		}
	}

	count := 0
	word := "balloon"

	for {
		for _, l := range word {
			if letters[l] == 0 {
				return count
			}
			letters[l]--
		}
		count++
	}
	return count
}

func MaxNumberOfBalloonsOneLoop(text string) int {
	count := make(map[rune]int)
	for _, c := range text {
		count[c]++
	}

	b := count['b']
	a := count['a']
	l := count['l'] / 2
	o := count['o'] / 2
	n := count['n']

	min := b
	if a < min {
		min = a
	}
	if l < min {
		min = l
	}
	if o < min {
		min = o
	}
	if n < min {
		min = n
	}

	return min
}
