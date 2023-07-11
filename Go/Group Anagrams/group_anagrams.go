package group_anagrams

import "sort"

func GroupAnagrams(strs []string) [][]string {
	r := make(map[string][]string)
	for _, w := range strs {
		c := []rune(w)
		sort.Slice(c, func(i int, j int) bool {
			return c[i] < c[j]
		})
		r[string(c)] = append(r[string(c)], w)
	}

	s := make([][]string, 0)
	for _, w := range r {
		s = append(s, w)
	}
	return s
}
