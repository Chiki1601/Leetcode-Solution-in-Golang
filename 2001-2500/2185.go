func prefixCount(words []string, pref string) int {
    count := 0
    
    for _, val := range words {
        pos := strings.Index(val, pref)
        if pos != -1 && pos == 0 {
            count++
        }
    }
    
    return count
}
