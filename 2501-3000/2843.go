func countSymmetricIntegers(low int, high int) int {
    cnt := 0
    for i := low; i <= high; i++ {
        s := strconv.Itoa(i)
        if len(s)%2 == 0 {
            mid := len(s) / 2
            sum1, sum2 := 0, 0
            for _, ch := range s[:mid] {
                sum1 += int(ch - '0')
            }
            for _, ch := range s[mid:] {
                sum2 += int(ch - '0')
            }
            if sum1 == sum2 {
                cnt++
            }
        }
    }
    return cnt
}
