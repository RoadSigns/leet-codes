package valid_parentheses

func IsValid(s string) bool {
	stack := make([]rune, 0, len(s))
	for _, r := range s {
		if isOpeningBracket(r) {
			stack = append(stack, r)
			continue
		}

		if !isClosingBracket(r) {
			continue
		}

		if len(stack) == 0 {
			return false
		}

		i := len(stack) - 1
		if isMatchingBracket(stack[i], r) {
			stack = stack[:i]
			continue
		}
		return false
	}

	if len(stack) != 0 {
		return false
	}

	return true
}

func isMatchingBracket(o, c rune) bool {
	if o == '[' && c == ']' {
		return true
	}

	if o == '{' && c == '}' {
		return true
	}

	if o == '(' && c == ')' {
		return true
	}

	return false
}

func isOpeningBracket(r rune) bool {
	if r == '{' {
		return true
	}

	if r == '[' {
		return true
	}

	if r == '(' {
		return true
	}

	return false
}

func isClosingBracket(r rune) bool {
	if r == '}' {
		return true
	}

	if r == ']' {
		return true
	}

	if r == ')' {
		return true
	}

	return false
}
