package jewels

func NumJewelsInStones(jewels string, stones string) int {
	j := make(map[rune]interface{})
	for _, jewel := range jewels {
		j[jewel] = nil
	}
	count := 0
	for _, stone := range stones {
		if _, ok := j[stone]; ok != false {
			count++
		}
	}
	return count
}
