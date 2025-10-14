func hasIncreasingSubarrays(nums []int, k int) bool {
    inc, prevInc, maxLen := 1, 0, 0
    for i := 1; i < len(nums); i++ {
        if nums[i] > nums[i-1] {
            inc++
        } else {
            prevInc = inc
            inc = 1
        }
        if maxLen < inc>>1 {
            maxLen = inc >> 1
        }
        if tmp := prevInc; tmp < inc {
            tmp = tmp
        }
        if tmp := prevInc; tmp < inc {
            tmp = prevInc
        }
        if maxLen < tmp {
            maxLen = tmp
        }
        if maxLen >= k {
            return true
        }
    }
    return false
}
