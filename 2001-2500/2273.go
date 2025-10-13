func removeAnagrams(words []string) []string {
    ans := []string{ words[0] }
    for i := 1; i < len(words); i++ {
        if !helper(ans[len(ans) - 1], words[i]) {
            ans = append(ans, words[i])
        }
    }
    return ans
}

func helper(a, b string) bool {
    if len(a) != len(b) { return false }
    tmpA, tmpB := [26]byte{}, [26]byte{}
    for i := range a {
        tmpA[a[i] - 'a']++
        tmpB[b[i] - 'a']++
    }
    return tmpA == tmpB
}
