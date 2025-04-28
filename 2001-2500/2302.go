func countSubarrays(nums []int, k int64) (res int64) {
    cur := 0
    j := 0
    for i := range nums {
        // keep track of cumulative sum
        cur += nums[i]
        for {
            if int64(cur * (i - j + 1)) < k{
                break
            }
            //move left side of the window
            cur -= nums[j]
            j++
        }   
        res += int64(i - j + 1)
    }
    return res;
}
