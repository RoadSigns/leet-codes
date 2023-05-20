package length_of_last_word

import (
	"strings"
)

func LengthOfLastWord(s string) int {
	slice := strings.Fields(s)
	return len(slice[len(slice)-1])
}
