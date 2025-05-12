func findEvenNumbers(digits []int) []int {
    ans := []int{}
    tmp := make(map[int]int)
    for _, d := range digits { tmp[d]++ }
    for i := 100; i < 999; i += 2 {
        curTmp := make(map[int]int)
        I := i
        for I > 0 {
            curTmp[I % 10]++
            I /= 10
        }
        b := true
        for i := 0; i < 10; i++ {
            if curTmp[i] > tmp[i] {
                b = false
                break
            }
        }
        if b { ans = append(ans, i) }
    }
    return ans
}
