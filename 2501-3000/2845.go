func countInterestingSubarrays(nums []int, modulo int, k int) int64 {
    cnt := map[int]int{0: 1}
    prefix, res := 0, int64(0)
    for _, num := range nums {
        if num % modulo == k {
            prefix++
        }
        key := (prefix - k + modulo) % modulo
        res += int64(cnt[key])
        cnt[prefix % modulo]++
    }
    return res
}
