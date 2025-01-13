func minimumLength(s string) int {
    charFrequency := make([]int, 26)
    totalLength := 0
    for _, char := range s {
        charFrequency[char - 'a']++
    }
    for _, frequency := range charFrequency {
        if frequency == 0 {
            continue
        }
        if frequency % 2 == 0 {
            totalLength += 2
        } else {
            totalLength += 1
        }
    }
    return totalLength
}
