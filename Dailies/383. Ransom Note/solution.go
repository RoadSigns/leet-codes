func canConstruct(ransomNote string, magazine string) bool {
    m := make(map[string]int)
    for _, c := range magazine {
        m[string(c)]++
    }
    
    for _, n := range ransomNote {
        if m[string(n)] == 0 {
            return false
        }
        m[string(n)]--
    }
    
    return true
}