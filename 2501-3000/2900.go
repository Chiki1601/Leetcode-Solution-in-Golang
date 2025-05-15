func getLongestSubsequence(words []string, groups []int) []string {
    build := func(start int) []string {
        res := []string{}
        expect := start
        for i := 0; i < len(words); i++ {
            if groups[i] == expect {
                res = append(res, words[i])
                expect ^= 1
            }
        }
        return res
    }
    res1, res2 := build(0), build(1)
    if len(res1) >= len(res2) {
        return res1
    }
    return res2
}
