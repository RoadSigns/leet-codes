package simplify_path

import "strings"

func SimplifyPath(path string) string {
	stack := make([]string, 0)
	d := make([]rune, 0)

	for _, c := range path {
		if c == '/' {
			if len(d) > 0 {
				stack = append(stack, string(d))
			}
			d = []rune{}
			continue
		}
		d = append(d, c)
	}

	if len(d) > 0 {
		stack = append(stack, string(d))
	}

	result := make([]string, 0)
	for _, s := range stack {
		if s == "." {
			continue
		}

		if s == ".." {
			if len(result) > 0 {
				result = result[:len(result)-1]
			}
			continue
		}

		result = append(result, s)
	}

	return "/" + strings.Join(result, "/")
}
