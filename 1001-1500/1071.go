func gcdOfStrings(str1 string, str2 string) string {
    divisors1 := getDivisor(str1)
    divisors2 := getDivisor(str2)
    longestDivisor := ""
    for i := range divisors1 {
        for j := range divisors2 {
            if divisors1[i] == divisors2[j] && len(divisors1[i]) > len(longestDivisor) {
                longestDivisor = divisors1[i]
            }
        }
    }
    return longestDivisor
}

func getDivisor(str string) []string {
    divisors := make([]string, 0)
    for l := 1; l <= len(str); l++ {
        if len(str) % l != 0 {
            continue
        }
        cand := str[:l]
        isDivisor := true
        for i := 0; i < len(str); i+=l {
            if str[i:i+l] != cand {
                isDivisor = false
                break
            }
        }
        if isDivisor {
            divisors = append(divisors, cand)
        }
    }
    return divisors
}
