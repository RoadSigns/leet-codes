package find_the_index_of_first_occurrence_in_a_string

func StrStr(haystack string, needle string) int {
	needleLength := len(needle)

	for i, _ := range haystack {
		if i+needleLength > len(haystack) {
			return -1
		}

		if haystack[i:i+needleLength] == needle {
			return i
		}
	}

	return -1
}
