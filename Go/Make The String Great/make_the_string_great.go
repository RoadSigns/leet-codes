package make_the_string_great

import "unicode"

func MakeGood(s string) string {
	stack := make([]rune, 0)
	for _, c := range s {
		if len(stack) > 0 {
			if unicode.IsUpper(c) && unicode.IsUpper(stack[len(stack)-1]) {
				stack = append(stack, c)
				continue
			}

			if c == unicode.ToUpper(stack[len(stack)-1]) {
				stack = stack[:len(stack)-1]
				continue
			}

			if unicode.ToUpper(c) == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
				continue
			}
		}
		stack = append(stack, c)
	}
	return string(stack)
}
