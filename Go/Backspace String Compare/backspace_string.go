package backspace_string

func BackspaceCompare(s string, t string) bool {
	return removeBackspace(s) == removeBackspace(t)
}

func removeBackspace(s string) string {
	stack := make([]rune, 0)

	for _, c := range s {
		if c != '#' {
			stack = append(stack, c)
			continue
		}

		if len(stack) > 0 {
			stack = stack[:len(stack)-1]
		}
	}

	return string(stack)
}
