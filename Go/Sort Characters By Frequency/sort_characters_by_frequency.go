package frequency

import "sort"

func FrequencySort(s string) string {
	c := make(map[rune]int)
	for _, l := range s {
		c[l]++
	}

	keys := make([]rune, 0, len(c))
	for k, _ := range c {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return c[keys[i]] > c[keys[j]]
	})

	res := make([]rune, 0, len(s))
	for _, k := range keys {
		for i := 0; i < c[k]; i++ {
			res = append(res, k)
		}
	}

	return string(res)
}
