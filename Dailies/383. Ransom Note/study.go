func canConstruct(ransomNote string, magazine string) bool {
	letters := [26]int{}
	for i := 0; i < len(magazine); i++ {
		letters[magazine[i]%97]++
	}
	for i := 0; i < len(ransomNote); i++ {
		if letters[ransomNote[i]%97] == 0 {
			return false
		}
		letters[ransomNote[i]%97]--
	}
	return true
}